#!/bin/bash

set -e

if [ -z "$1" ]; then
    echo "required patch/minor/major" 1>&2
    exit 1
fi

new_version=$(gobump "$1" -w -v | jq -r '.[]')

git add ./*.go
git commit -m "Bump version $new_version"
git push origin master
make cross
ghr --username yuuki1 --replace --draft "Release v$new_version" snapshot/
