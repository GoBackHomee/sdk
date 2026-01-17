# Gobackhomee SDK

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**The Bridge** - Shared types and client libraries for the Gobackhomee ecosystem.

## Overview

The SDK provides MIT-licensed shared Go structs and client libraries, allowing proprietary applications to interface with the Gobackhomee Core engine without license infection.

## Installation

```bash
go get github.com/gobackhomee/sdk
```

## Features

### 🔐 Web3-Native Authentication
- SIWE (Sign-In with Ethereum) support
- Wallet-first identity (not email-first)
- Multi-chain support (Ethereum, Solana, Polygon)

### 🚀 Vercel-Style Deployments
- Project and deployment management
- Zero-downtime updates
- Content-addressed deployments (Merkle roots)

### 🤖 AI Integration
- Schema generation from natural language
- Vector embeddings for RAG applications
- Local-first (Ollama) with cloud fallback

## Quick Start

```go
package main

import (
    "context"
    "log"
    
    "github.com/gobackhomee/sdk/client"
)

func main() {
    // Create client with Web3 auth
    c := client.New("http://localhost:8080", 
        client.WithWalletAuth(&client.WalletAuth{
            WalletAddress: "0x...",
            SignMessage: yourSigningFunction,
        }),
    )
    
    // Create a project
    project, err := c.Projects().Create(context.Background(), "my-app")
    if err != nil {
        log.Fatal(err)
    }
    
    // Use AI to generate a schema
    schema, err := c.AI().GenerateSchema(context.Background(), 
        "User table with wallet address, email, and created timestamp")
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Project: %s, Schema: %s", project.Name, schema)
}
```

## Package Structure

```
sdk/
├── types/       # Core domain types (Identity, User, Project, Deployment)
├── config/      # Configuration structs for Core and Fleet
├── client/      # Go SDK client
└── pkg/         # Public utilities
```

## License

MIT - See [LICENSE](LICENSE) for details.
