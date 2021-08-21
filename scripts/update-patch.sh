#!/bin/bash

set -e
set +x

SCRIPTS_DIR="$(dirname "$0")"

pushd "$SCRIPTS_DIR/../playwright"
SCRIPTS_DIR="$(dirname "$0")"
echo "Creating patch..."
git add .
git diff --staged --full-index playwright-head > ../patches/main.patch

cd ..
git submodule update --init

popd
