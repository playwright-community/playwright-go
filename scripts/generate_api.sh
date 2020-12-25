#!/bin/bash

set -e
set +x

echo "Generating types"
node scripts/generate-stucts.js > types.go
echo "Generated types"

echo "Formatting types"
go fmt types.go
echo "Formatted types"
