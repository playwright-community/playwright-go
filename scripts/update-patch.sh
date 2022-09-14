#!/bin/bash

set -e
set +x

SCRIPTS_DIR="$(dirname "$0")"

pushd "$SCRIPTS_DIR/../playwright"
SCRIPTS_DIR="$(dirname "$0")"
echo "Creating patch..."
git add .
git diff --full-index 396487f..main > ../patches/main.patch
git reset --hard playwright-build^1
cd ..

popd
