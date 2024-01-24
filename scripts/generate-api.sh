#!/bin/bash

set -e
set +x
PWD=$(pwd -P)

"$PWD"/scripts/apply-patch.sh

echo
echo "Generating code"
echo "=================="
PLAYWRIGHT_DIR="playwright"
node $PLAYWRIGHT_DIR/utils/doclint/generateGoApi.js
errCode=$?
if [ $errCode -ne 0 ]; then
  echo 
  exit $errCode
else
  echo "Done!"
  echo 
fi
mv $PLAYWRIGHT_DIR/utils/doclint/generate_types/go/generated-{enums,interfaces,structs}.go .
# fmt first or not, gofumpt's result will be different
go fmt generated-{enums,interfaces,structs}.go > /dev/null
gofumpt -w generated-{enums,interfaces,structs}.go > /dev/null

echo "Updating README"
echo "==============="
go run scripts/install-browsers/main.go
go run scripts/update-readme-versions/main.go

git submodule update
