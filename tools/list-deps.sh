#!/bin/sh
# Find and print the dependencies of all go.mod files under the current
# directory. Run `go mod vendor` before running to get versions of all
# dependencies.
set -eu

cat $(find . -name 'go.mod' -type f) |
grep '	' |
tr -d '\t' |
cut -d ' ' -f 1 |
sort |
uniq
