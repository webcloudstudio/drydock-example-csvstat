#!/bin/sh
# Runs go vet and go test from the application root, whatever the caller's working directory.
# Exits non-zero at the first failing command, so the first failure is never masked.
set -e

ROOT=$(cd "$(dirname "$0")/.." && pwd)
cd "$ROOT"

go vet ./...
go test ./...
