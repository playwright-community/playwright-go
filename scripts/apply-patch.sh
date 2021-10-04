#!/bin/bash

set -e

BRANCH_NAME_BUILD="playwright-build"
SCRIPTS_DIR="$(dirname "$0")"

echo "Applying patches..."

pushd "$SCRIPTS_DIR/.."
git submodule update --init
mkdir -p tests/assets
cp -r playwright/tests/assets/* tests/assets/
pushd playwright

git checkout HEAD --detach

if git show-ref -q --heads "$BRANCH_NAME_BUILD"; then
  git branch -D "$BRANCH_NAME_BUILD"
fi

git checkout -b "$BRANCH_NAME_BUILD"
git apply --whitespace=nowarn ../patches/*
git add -A
git commit -m "Applied patches"

popd