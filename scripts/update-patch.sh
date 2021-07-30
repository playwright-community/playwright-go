#!/bin/bash

echo "Creating patch..."

cd "$(dirname "$0")"

cd ../playwright 

git diff --full-index --src-prefix="a/playwright/" --dst-prefix="b/playwright/" > ../patches/main.patch $(git rev-parse HEAD^1)..$(git rev-parse HEAD)

git reset --hard $(git rev-parse HEAD^1)

cd -