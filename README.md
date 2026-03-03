# VaultFlow 🔐

Secure Secrets Management CLI - AES-256-GCM encrypted local vault for developers.

## Overview

VaultFlow is a production-ready CLI tool for managing secrets locally with military-grade encryption. Perfect for developers who need a simple, secure way to store API keys, passwords, and other sensitive data.

## Features

- **AES-256-GCM Encryption** - Military-grade encryption
- **Local Storage** - Secrets never leave your machine
- **Simple CLI** - Easy to use commands
- **Zero Dependencies** - Static binary
- **Cross-Platform** - Linux, macOS, Windows

## Installation

### From Binary
```bash
# Linux
curl -sL https://github.com/dablon/vaultflow/releases/download/v1.0.0/vaultflow-linux-amd64 -o vaultflow
chmod +x vaultflow

# macOS
curl -sL https://github.com/dablon/vaultflow/releases/download/v1.0.0/vaultflow-darwin-amd64 -o vaultflow

# Windows
curl -sL https://github.com/dablon/vaultflow/releases/download/v1.0.0/vaultflow-windows-amd64.exe -o vaultflow.exe
```

### From Source
```bash
go install github.com/dablon/vaultflow@latest
```

## Usage

### Store a Secret
```bash
vaultflow set db_password "my-secret-password"
# ✓ Secret stored securely
```

### Retrieve a Secret
```bash
vaultflow get db_password
# my-secret-password
```

### List All Secrets
```bash
vaultflow list
# Stored secrets:
#   • db_password
#   • api_key
```

### Delete a Secret
```bash
vaultflow delete db_password
# ✓ Secret deleted
```

## Configuration

### Environment Variables
- `VAULT_MASTER_KEY` - Master encryption key (optional, generates default)

### Default Vault Location
- Linux/macOS: `~/.vaultflow/secrets.json`
- Windows: `%USERPROFILE%\.vaultflow\secrets.json`

## Security

### Encryption Details
- **Algorithm**: AES-256-GCM
- **Key Derivation**: Direct 32-byte key (padded if shorter)
- **Nonce**: Random 12-byte nonce per encryption

### Security Best Practices
1. Set `VAULT_MASTER_KEY` environment variable
2. Don't commit `secrets.json` to version control
3. Use different keys for different environments

## Docker

```bash
docker run --rm -v ~/.vaultflow:/root/.vaultflow \
  -e VAULT_MASTER_KEY=your-key \
  vaultflow list
```

## Architecture

```
┌─────────────┐     ┌──────────────┐     ┌─────────────┐
│   CLI       │────▶│   Vault      │────▶│   Crypto    │
│  (Cobra)    │     │   Manager    │     │  (AES-256)  │
└─────────────┘     └──────────────┘     └─────────────┘
                           │
                           ▼
                    ┌─────────────┐
                    │   JSON      │
                    │   File      │
                    └─────────────┘
```

## License

MIT License - See LICENSE file for details.
