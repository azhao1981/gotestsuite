#!/bin/bash

version=$1

if [ -z "$version" ]; then
    echo "version is required"
    exit 1
fi

git tag v$version
git push origin v$version
go list -m github.com/azhao1981/gotestsuite@v$version
