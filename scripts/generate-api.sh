#!/bin/bash

set -e
set +x
PWD=$(pwd -P)

$PWD/scripts/apply-patch.sh

go run scripts/install-browsers/main.go

echo "Generating Interfaces"
node scripts/generate-interfaces.js > generated-interfaces.go
go fmt generated-interfaces.go > /dev/null
echo "Generated Interfaces"

PLAYWRIGHT_DIR="playwright"

node $PLAYWRIGHT_DIR/utils/doclint/generateGoApi.js
mv $PLAYWRIGHT_DIR/utils/doclint/generate_types/go/generated-{enums,structs}.go .
go fmt generated-{enums,structs}.go > /dev/null

echo "Updating README"
go run scripts/update-readme-versions/main.go

git submodule update

