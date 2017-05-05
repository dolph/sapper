#!/bin/bash
set -ex

HOSTNAME=$1

response=$(curl --verbose https://$HOSTNAME/ 2> /tmp/smoke-stderr)
response_stderr=$(</tmp/smoke-stderr)
rm /tmp/smoke-stderr

# Test that the response contains something that looks like an IP.
if [[ ! $response =~ ^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$ ]]; then
    echo "Unexpected response body:" $response
    exit 1
fi

echo $response_stderr | grep "text/plain"
