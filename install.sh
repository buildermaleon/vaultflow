#!/bin/bash
set -e
mkdir -p ~/.local/bin
curl -sL "https://github.com/dablon/vaultflow/releases/download/v1.0.0/vaultflow-linux-amd64" -o ~/.local/bin/vaultflow
chmod +x ~/.local/bin/vaultflow
