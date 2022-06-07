#!/bin/bash

set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd ${DIR}/.. || exit 1

mkdir -p dist

go build \
  -trimpath \
  -ldflags "-s -w" \
  -o dist/${2} \
  ${1}

if [[ $(which upx) != "" ]]; then
  upx dist/${2}
fi
