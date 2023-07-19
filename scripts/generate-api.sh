#!/bin/bash

set -e
set +x
PWD=$(pwd -P)

"$PWD"/scripts/apply-patch.sh

echo "Generating Interfaces"
echo "====================="
node scripts/generate-interfaces.js > generated-interfaces.go
go fmt generated-interfaces.go > /dev/null

echo "Generating structs"
echo "=================="
PLAYWRIGHT_DIR="playwright"
node $PLAYWRIGHT_DIR/utils/doclint/generateGoApi.js
mv $PLAYWRIGHT_DIR/utils/doclint/generate_types/go/generated-{enums,structs}.go .
go fmt generated-{enums,structs}.go > /dev/null

echo "Validating Interfaces"
echo "====================="
node scripts/validate-interfaces.js
errCode=$?
if [ $errCode -ne 0 ]; then
  echo 
  echo "Please make sure to implement the necessary interface methods."
  exit $errCode
else
  echo "Done!"
  echo 
fi

echo "Updating README"
echo "==============="
go run scripts/install-browsers/main.go
go run scripts/update-readme-versions/main.go

git submodule update
