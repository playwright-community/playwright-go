#!/bin/bash

set -e
set +x

echo "Generating types"
node scripts/generate-structs.js > generated_types.go
echo "Generated types"

echo "Formatting types"
go fmt generated_types.go
echo "Formatted types"
