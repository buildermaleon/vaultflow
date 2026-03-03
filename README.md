# VaultFlow

Secure secrets management with AES-256-GCM encryption.

## Installation

```bash
go install github.com/dablon/vaultflow@latest
```

Or build from source:

```bash
git clone https://github.com/dablon/vaultflow.git
cd vaultflow
go install ./cmd
```

## Usage

```bash
# Store a secret
vaultflow set db_password mysecretpassword

# Retrieve a secret
vaultflow get db_password

# List all secrets
vaultflow list

# Delete a secret
vaultflow delete db_password
```

## Features

- AES-256-GCM encryption
- Master key protection via environment variable
- Local encrypted storage
- Simple CLI interface

## Environment Variables

- `VAULT_MASTER_KEY` - Master encryption key

## License

MIT
