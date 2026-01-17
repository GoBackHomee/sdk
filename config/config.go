// Copyright (C) 2024 Gobackhomee
// SPDX-License-Identifier: MIT

// Package config provides configuration types for Gobackhomee services.
package config

import "time"

// CoreConfig defines the configuration for a Core engine instance
type CoreConfig struct {
	// Server configuration
	Server ServerConfig `json:"server" yaml:"server"`

	// Database configuration (Postgres-first like Supabase)
	Database DatabaseConfig `json:"database" yaml:"database"`

	// Web3 blockchain configuration
	Web3 Web3Config `json:"web3" yaml:"web3"`

	// AI configuration (Ollama-first)
	AI AIConfig `json:"ai" yaml:"ai"`

	// Hosting configuration for Vercel-style deployments
	Hosting HostingConfig `json:"hosting" yaml:"hosting"`
}

// ServerConfig for HTTP/gRPC server settings
type ServerConfig struct {
	Host         string        `json:"host" yaml:"host"`
	Port         int           `json:"port" yaml:"port"`
	ReadTimeout  time.Duration `json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout" yaml:"write_timeout"`
	TLSCert      string        `json:"tls_cert,omitempty" yaml:"tls_cert"`
	TLSKey       string        `json:"tls_key,omitempty" yaml:"tls_key"`
}

// DatabaseConfig prioritizes Postgres (Supabase-style) with fallbacks
type DatabaseConfig struct {
	// Driver: postgres (default), sqlite, badgerdb
	Driver string `json:"driver" yaml:"driver"`

	// DSN connection string
	DSN string `json:"dsn" yaml:"dsn"`

	// Pool configuration
	MaxOpenConns int `json:"max_open_conns" yaml:"max_open_conns"`
	MaxIdleConns int `json:"max_idle_conns" yaml:"max_idle_conns"`

	// Enable pgvector for AI embeddings
	EnableVector bool `json:"enable_vector" yaml:"enable_vector"`
}

// Web3Config for blockchain integrations
type Web3Config struct {
	// Enabled chains
	EnabledChains []ChainConfig `json:"enabled_chains" yaml:"enabled_chains"`

	// SIWE (Sign-In with Ethereum) settings
	SIWE SIWEConfig `json:"siwe" yaml:"siwe"`
}

// ChainConfig for individual blockchain configuration
type ChainConfig struct {
	// Name: ethereum, solana, polygon, etc.
	Name string `json:"name" yaml:"name"`

	// RPC endpoint URL
	RPCURL string `json:"rpc_url" yaml:"rpc_url"`

	// Chain ID for EVM chains
	ChainID int `json:"chain_id,omitempty" yaml:"chain_id"`

	// WebSocket URL for real-time events
	WSURL string `json:"ws_url,omitempty" yaml:"ws_url"`
}

// SIWEConfig for Sign-In with Ethereum
type SIWEConfig struct {
	Enabled    bool   `json:"enabled" yaml:"enabled"`
	Domain     string `json:"domain" yaml:"domain"`
	Statement  string `json:"statement" yaml:"statement"`
	SessionTTL string `json:"session_ttl" yaml:"session_ttl"`
}

// AIConfig for AI/LLM integrations (local-first with Ollama)
type AIConfig struct {
	// Provider: ollama (default), openai, anthropic
	Provider string `json:"provider" yaml:"provider"`

	// Endpoint for local Ollama or API URL
	Endpoint string `json:"endpoint" yaml:"endpoint"`

	// APIKey for cloud providers
	APIKey string `json:"api_key,omitempty" yaml:"api_key"`

	// DefaultModel for inference
	DefaultModel string `json:"default_model" yaml:"default_model"`

	// EmbeddingModel for vector operations
	EmbeddingModel string `json:"embedding_model" yaml:"embedding_model"`
}

// HostingConfig for Vercel-style static hosting
type HostingConfig struct {
	// DataDir where sites are stored
	DataDir string `json:"data_dir" yaml:"data_dir"`

	// MaxUploadSize in bytes
	MaxUploadSize int64 `json:"max_upload_size" yaml:"max_upload_size"`

	// EnableSRI for Subresource Integrity verification
	EnableSRI bool `json:"enable_sri" yaml:"enable_sri"`

	// SPAFallback enables SPA routing (fallback to index.html)
	SPAFallback bool `json:"spa_fallback" yaml:"spa_fallback"`
}

// FleetConfig for multi-node orchestration
type FleetConfig struct {
	// Master node endpoint
	MasterEndpoint string `json:"master_endpoint" yaml:"master_endpoint"`

	// Node registration settings
	NodeID   string `json:"node_id" yaml:"node_id"`
	NodeName string `json:"node_name" yaml:"node_name"`

	// RBAC settings
	RBAC RBACConfig `json:"rbac" yaml:"rbac"`
}

// RBACConfig for role-based access control
type RBACConfig struct {
	Enabled     bool     `json:"enabled" yaml:"enabled"`
	AdminRoles  []string `json:"admin_roles" yaml:"admin_roles"`
	DefaultRole string   `json:"default_role" yaml:"default_role"`
}
