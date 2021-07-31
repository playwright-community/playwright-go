#!/bin/bash

set -e
set +x
PWD=$(pwd -P)

$PWD/scripts/apply-patch.sh

echo "Generating Interfaces"
node scripts/generate-interfaces.js > generated_interfaces.go
go fmt generated_interfaces.go > /dev/null
echo "Generated Interfaces"

PLAYWRIGHT_DIR="playwright"

node $PLAYWRIGHT_DIR/utils/doclint/generateGoApi.js
mv $PLAYWRIGHT_DIR/utils/doclint/generate_types/go/generated-{enums,structs}.go .
rm $PLAYWRIGHT_DIR/utils/doclint/generate_types/go/generated-interfaces.go
go fmt generated-{enums,structs}.go > /dev/null

# echo "Validating API"
# node scripts/validate-interfaces.js
# echo "Validated API"

# echo "Updating README"
go run scripts/update-readme-versions/main.go

pushd $PLAYWRIGHT_DIR

git add -A

git reset --hard $(git rev-parse HEAD^1)

popd