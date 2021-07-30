#!/bin/bash

echo "Applying patches..."

cd "$(dirname "$0")"

cd ..

git apply --whitespace=nowarn patches/*

pushd playwright

git add -A

git commit -m "Applied patches"

popd

cd -