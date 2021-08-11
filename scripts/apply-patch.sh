#!/bin/bash

set -e

BRANCH_NAME="playwright-build"
SCRIPTS_DIR="$(dirname "$0")"

echo "Applying patches..."

pushd "$SCRIPTS_DIR/.."
pushd playwright

git reset --hard HEAD

if git show-ref -q --heads "$BRANCH_NAME"; then
  git branch -D "$BRANCH_NAME"
fi

git apply --whitespace=nowarn ../patches/*
git add -A
git checkout -b "$BRANCH_NAME"
git commit -m "Applied patches"

popd
popd