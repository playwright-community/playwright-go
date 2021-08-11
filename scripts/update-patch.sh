#!/bin/bash

set -e
set +x

SCRIPTS_DIR="$(dirname "$0")"

pushd "$SCRIPTS_DIR/../playwright"
SCRIPTS_DIR="$(dirname "$0")"
echo "Creating patch..."

git diff --full-index HEAD^1..HEAD > ../patches/main.patch
git reset --hard HEAD^1
git branch -D playwright-build

popd
