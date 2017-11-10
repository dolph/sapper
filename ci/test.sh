#!/bin/bash
set -ex

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $DIR/../

# Install requirements.
go get "gopkg.in/urfave/cli.v1"

# Run unit tests.
go test
