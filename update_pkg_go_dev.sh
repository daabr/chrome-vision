#!/bin/bash
#
# Trigger an update of https://pkg.go.dev/github.com/daabr/chrome-vision
# based on the instructions in https://go.dev/about#adding-a-package.

GOPROXY=https://proxy.golang.org GO111MODULE=on go get github.com/daabr/chrome-vision@v0.0.0
