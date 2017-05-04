#!/bin/bash
set -ex

HOSTNAME=$1

response=`curl https://$HOSTNAME/`

# Test that the response contains something that looks like an IP.
if [[ $response =~ ^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$ ]]; then
    exit 0
else
    echo "Received:" $response
    exit 1
fi
