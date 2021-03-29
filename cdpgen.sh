#!/bin/bash
#
# Generate the source code of the Go bindings for all the commands, events
# and types in the Chrome DevTools Protocol (CDP) - from the latest Chromium
# "tip-of-tree" (tot) definitions (see the details and API documentation in
# in https://chromedevtools.github.io/devtools-protocol/tot, mirrored in
# https://github.com/ChromeDevTools/devtools-protocol).

go generate ./cmd/cdpgen

for f in ./pkg/cdp/**/*.go; do
    ~/go/bin/goimports $f > ./tmp.go
    mv -f ./tmp.go $f
done
