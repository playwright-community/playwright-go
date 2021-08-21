#!/bin/bash

set -e

BRANCH_NAME_HEAD="playwright-head"
BRANCH_NAME_BUILD="playwright-build"
SCRIPTS_DIR="$(dirname "$0")"

echo "Applying patches..."

pushd "$SCRIPTS_DIR/.."
git submodule update --init
pushd playwright

git reset --hard HEAD
git checkout master

if git show-ref -q --heads "$BRANCH_NAME_HEAD"; then
  git branch -D "$BRANCH_NAME_HEAD"
fi
if git show-ref -q --heads "$BRANCH_NAME_BUILD"; then
  git branch -D "$BRANCH_NAME_BUILD"
fi

git checkout -b "$BRANCH_NAME_HEAD"
git checkout -b "$BRANCH_NAME_BUILD"
git apply --whitespace=nowarn ../patches/*
git add -A
git commit -m "Applied patches"

popd
