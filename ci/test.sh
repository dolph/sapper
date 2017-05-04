#!/bin/bash
set -ex

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Run unit tests.
cd $DIR/../src/skeleton/
go test
