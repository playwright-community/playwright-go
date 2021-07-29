#!/bin/bash

echo "Creating patch..."

cd "$(dirname "$0")"

cd ../playwright 

git diff --diff-algorithm=myers --full-index --staged > ../patches/main.patch

git reset --hard

cd - 