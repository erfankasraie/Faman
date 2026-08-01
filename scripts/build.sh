#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

echo "Building faman..."
go build -ldflags="-s -w" -o faman ./cmd/faman
echo "Done: ./faman"
./faman version
