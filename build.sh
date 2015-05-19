#!/bin/bash

docker run --rm -e CGO_ENABLED=0 -v "$PWD":/usr/src/redirect -w /usr/src/redirect golang go build -v -a -installsuffix cgo -ldflags '-s'