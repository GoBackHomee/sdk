// Copyright (C) 2024 Gobackhomee
// SPDX-License-Identifier: MIT

// Package types defines core domain types shared across the Gobackhomee ecosystem.
// These types are intentionally MIT-licensed to allow proprietary applications
// to interface with the Core engine without license infection.
package types

import (
	"time"
)

// Identity represents a Web3-native identity rooted in wallet addresses.
// Unlike traditional auth systems that use emails, we prioritize wallet-first identity.
type Identity struct {
	// ID is the unique identifier (UUID)
	ID string `json:"id"`

	// WalletAddress is the primary identifier - Web3 native
	WalletAddress string `json:"wallet_address"`

	// Chain identifies which blockchain (ethereum, solana, etc.)
	Chain string `json:"chain"`

	// PublicKey for signature verification
	PublicKey string `json:"public_key,omitempty"`

	// Email is optional, not required for Web3 auth
	Email string `json:"email,omitempty"`

	// Metadata for extensibility
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Timestamps
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// User extends Identity with application-level data
type User struct {
	Identity

	// DisplayName for UI purposes
	DisplayName string `json:"display_name,omitempty"`

	// Avatar URL
	AvatarURL string `json:"avatar_url,omitempty"`

	// Roles for RBAC
	Roles []string `json:"roles,omitempty"`

	// Active status
	Active bool `json:"active"`
}

// Project represents a deployment project (like Vercel projects)
type Project struct {
	// ID is the unique project identifier
	ID string `json:"id"`

	// Name is the human-readable project name
	Name string `json:"name"`

	// OwnerID links to the owning Identity
	OwnerID string `json:"owner_id"`

	// Domain configuration
	Domains []string `json:"domains,omitempty"`

	// CurrentVersion is the active deployment hash
	CurrentVersion string `json:"current_version,omitempty"`

	// Framework detection (react, vue, static, etc.)
	Framework string `json:"framework,omitempty"`

	// Timestamps
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Deployment represents a single deployment version
type Deployment struct {
	// ID is unique deployment identifier
	ID string `json:"id"`

	// ProjectID links to parent project
	ProjectID string `json:"project_id"`

	// Version is the semantic version or commit hash
	Version string `json:"version"`

	// Hash is the content-addressed Merkle root for trustless verification
	Hash string `json:"hash"`

	// Status: pending, building, ready, failed
	Status string `json:"status"`

	// URL where this deployment is accessible
	URL string `json:"url,omitempty"`

	// Timestamps
	CreatedAt time.Time  `json:"created_at"`
	ReadyAt   *time.Time `json:"ready_at,omitempty"`
}
