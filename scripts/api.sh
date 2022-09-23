#!/bin/sh

figlet Api
reflex -r '(\.go|go\.mod)' -s go run ./cmd/api/main.go
