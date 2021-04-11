// The cdpgen program generates the source code of the Go bindings for all the
// commands, events and types in the Chrome DevTools Protocol (CDP) - from the
// latest Chromium "tip-of-tree" (tot) definitions (see the details and API
// documentation in https://chromedevtools.github.io/devtools-protocol/tot,
// mirrored in https://github.com/ChromeDevTools/devtools-protocol).
package main

//go:generate curl -sS https://raw.githubusercontent.com/ChromeDevTools/devtools-protocol/master/json/browser_protocol.json -o browser_protocol.json
//go:generate curl -sS https://raw.githubusercontent.com/ChromeDevTools/devtools-protocol/master/json/js_protocol.json -o js_protocol.json
//go:generate go run main.go
//go:generate rm browser_protocol.json
//go:generate rm js_protocol.json

import (
	"encoding/json"
	"log"
	"os"

	"github.com/daabr/chrome-vision/internal/cdpgen"
)

func main() {
	files := [...]string{"browser_protocol.json", "js_protocol.json"}
	for _, f := range files {
		bytes, err := os.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}
		var p cdpgen.Protocol
		json.Unmarshal([]byte(bytes), &p)
		cdpgen.Generate(p)
	}
}
