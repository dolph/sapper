#!/bin/bash
set -ex

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Run unit tests.
cd $DIR/../src/skeleton/

# Linting.
files_needing_linting=$(gofmt -l .)
if [[ $(gofmt -l .) ]]; then
    echo $files_needing_linting
    echo "Resolve differences with go fmt."
    exit 1
fi

go test
