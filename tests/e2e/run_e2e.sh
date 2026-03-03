#!/bin/bash
set -e

echo "=== VaultFlow E2E Test ==="

echo "1. Building binary..."
go build -o vaultflow ./cmd

echo "2. Running unit tests..."
go test -v -cover ./internal/...

echo "3. Testing set/get/list/delete..."
./vaultflow set test_key "test_value"
./vaultflow get test_key
./vaultflow list
./vaultflow delete test_key
./vaultflow list

echo "=== E2E Tests PASSED ==="
