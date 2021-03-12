#!/bin/bash

set -e
set +x


echo "Generating Interfaces"
node scripts/generate-interfaces.js > generated_interfaces.go
go fmt generated_interfaces.go > /dev/null
echo "Generated Interfaces"

PLAYWRIGHT_DIR="../playwright"

node $PLAYWRIGHT_DIR/utils/doclint/generateGoApi.js
cp $PLAYWRIGHT_DIR/utils/doclint/generate_types/go/generated-{enums,structs}.go .
go fmt generated-{enums,structs}.go > /dev/null

# echo "Validating API"
# node scripts/validate-interfaces.js
# echo "Validated API"

# echo "Updating README"
go run scripts/update-readme-versions/main.go
