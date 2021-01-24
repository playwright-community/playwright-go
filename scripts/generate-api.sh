#!/bin/bash

set -e
set +x

echo "Generating types"
node scripts/generate-structs.js > generated_types.go
go fmt generated_types.go > /dev/null
echo "Generated types"

echo "Generating Interfaces"
node scripts/generate-interfaces.js > generated_interfaces.go
go fmt generated_interfaces.go > /dev/null
echo "Generated Interfaces"

echo "Validating API"
node scripts/validate-interfaces.js || true
echo "Validated API"

echo "Updating README"
go run scripts/update-readme-versions/main.go
