#!/bin/bash

set -e

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# dep ensure

version=$(awk '/const Binary/ {print $NF}' < $DIR/internal/version/binary.go | sed 's/"//g')

for os in linux darwin windows; do 
    echo "... building v$version for $os/$arch"
done
