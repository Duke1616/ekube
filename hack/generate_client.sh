#!/bin/bash

set -e

GV="$1"

rm -rf ./pkg/client
./hack/generate_group.sh "client,lister,informer" ekube/pkg/client ekube/api "${GV}" --output-base=./  -h "$PWD/hack/boilerplate.go.txt"
mv ekube/pkg/client ./pkg/
rm -rf ./ekube
