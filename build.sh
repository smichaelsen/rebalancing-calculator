#!/usr/bin/env bash

GOOS=windows GOARCH=amd64 go build -o build/rebalance-win.exe ./cmd/rebalance/main.go
GOOS=darwin GOARCH=amd64 go build -o build/rebalance-intel ./cmd/rebalance/main.go
GOOS=darwin GOARCH=arm64 go build -o build/rebalance-silicon ./cmd/rebalance/main.go
