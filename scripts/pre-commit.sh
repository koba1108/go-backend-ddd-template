#!/bin/sh

echo Running static check for golang
go install honnef.co/go/tools/cmd/staticcheck@2022.1.2
staticcheck ./internals/...
