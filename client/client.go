// Copyright (C) 2024 Gobackhomee
// SPDX-License-Identifier: MIT

// Package client provides a Go client for interacting with Gobackhomee Core.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gobackhomee/sdk/types"
)

// Client is the Gobackhomee SDK client
type Client struct {
	baseURL    string
	httpClient *http.Client
	apiKey     string
	walletAuth *WalletAuth
}

// WalletAuth holds Web3 authentication credentials
type WalletAuth struct {
	WalletAddress string
	SignMessage   func(message string) (signature string, err error)
}

// Option configures the Client
type Option func(*Client)

// WithAPIKey sets an API key for authentication
func WithAPIKey(key string) Option {
	return func(c *Client) {
		c.apiKey = key
	}
}

// WithWalletAuth sets Web3 wallet authentication
func WithWalletAuth(auth *WalletAuth) Option {
	return func(c *Client) {
		c.walletAuth = auth
	}
}

// WithTimeout sets the HTTP client timeout
func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = d
	}
}

// New creates a new Gobackhomee client
func New(baseURL string, opts ...Option) *Client {
	c := &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Auth returns the authentication service
func (c *Client) Auth() *AuthService {
	return &AuthService{client: c}
}

// Projects returns the projects service
func (c *Client) Projects() *ProjectsService {
	return &ProjectsService{client: c}
}

// AI returns the AI service
func (c *Client) AI() *AIService {
	return &AIService{client: c}
}

// doRequest performs an HTTP request with authentication
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	return c.httpClient.Do(req)
}

// AuthService handles authentication operations (Web3-native)
type AuthService struct {
	client *Client
}

// SignInWithEthereum performs SIWE authentication
func (a *AuthService) SignInWithEthereum(ctx context.Context, message, signature string) (*types.Identity, error) {
	resp, err := a.client.doRequest(ctx, "POST", "/api/auth/siwe", map[string]string{
		"message":   message,
		"signature": signature,
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var identity types.Identity
	if err := json.NewDecoder(resp.Body).Decode(&identity); err != nil {
		return nil, err
	}

	return &identity, nil
}

// ProjectsService handles project operations
type ProjectsService struct {
	client *Client
}

// Create creates a new project
func (p *ProjectsService) Create(ctx context.Context, name string) (*types.Project, error) {
	resp, err := p.client.doRequest(ctx, "POST", "/api/projects", map[string]string{
		"name": name,
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var project types.Project
	if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
		return nil, err
	}

	return &project, nil
}

// List lists all projects for the authenticated user
func (p *ProjectsService) List(ctx context.Context) ([]types.Project, error) {
	resp, err := p.client.doRequest(ctx, "GET", "/api/projects", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var projects []types.Project
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		return nil, err
	}

	return projects, nil
}

// AIService handles AI operations
type AIService struct {
	client *Client
}

// GenerateSchema uses AI to generate a schema from natural language
func (a *AIService) GenerateSchema(ctx context.Context, description string) (string, error) {
	resp, err := a.client.doRequest(ctx, "POST", "/api/ai/schema", map[string]string{
		"description": description,
	})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Schema string `json:"schema"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Schema, nil
}

// Embed generates embeddings for the given text (for RAG applications)
func (a *AIService) Embed(ctx context.Context, text string) ([]float32, error) {
	resp, err := a.client.doRequest(ctx, "POST", "/api/ai/embed", map[string]string{
		"text": text,
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Embedding []float32 `json:"embedding"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Embedding, nil
}
