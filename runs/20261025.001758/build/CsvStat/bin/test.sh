#!/bin/sh
# CommandCenter Operation
# Name: Test
# Type: batch

# Runs the csvstat verification gate: go vet, then go test, across the module.
set -eu

PROJECT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
cd "$PROJECT_DIR"

go vet ./...
go test ./...
