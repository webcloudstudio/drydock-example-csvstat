#!/bin/sh
# CommandCenter Operation
# Name: Test
# Type: batch
#
# Runs go vet and the Go test suite from the application root.
set -eu

cd "$(dirname "$0")/.."
go vet ./...
go test ./...
