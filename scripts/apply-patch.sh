#!/bin/bash

echo "Applying patches..."

cd "$(dirname "$0")"

cd ../playwright

git apply --index --whitespace=nowarn ../patches/*

git add -A

cd -