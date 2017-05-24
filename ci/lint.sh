#!/bin/bash
set -ex

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $DIR/../

# Linting.
files_needing_linting=$(gofmt -l .)
if [[ $(gofmt -d .) ]]; then
    echo $files_needing_linting
    echo "Resolve differences with go fmt."
    exit 1
fi
