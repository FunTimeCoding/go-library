#!/bin/sh -e

gotestsum --format standard-quiet -- ./...
golangci-lint run
