#!/bin/sh
# CommandCenter Operation
# Name: Run Test Suite
# Type: batch

# Gates the project: formatting, vet, then the Go test suite. Exits non-zero on any failure.
set -eu

SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)
APP_DIR=$(cd "$SCRIPT_DIR/.." && pwd)
cd "$APP_DIR"

unformatted=$(gofmt -l .)
if [ -n "$unformatted" ]; then
    echo "gofmt needed for:" >&2
    echo "$unformatted" >&2
    exit 1
fi

go vet ./...
go test -race -count=1 ./...
