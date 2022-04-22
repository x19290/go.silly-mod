#!/bin/sh

set -eux

rm -f go.sum
go mod tidy
cp go.mod go.mod-kept
go test ./...
sed -i '6,$d' go.mod
