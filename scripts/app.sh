#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

if [ ! $(command -v leo) ]
then
	go install github.com/go-leo/leo/v3/cmd/leo@latest
	leo --version
fi