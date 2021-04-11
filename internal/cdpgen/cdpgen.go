// Package cdpgen generates the source code of the Go bindings for all the
// commands, events and types in the Chrome DevTools Protocol (CDP) - from the
// latest Chromium "tip-of-tree" (tot) definitions (see the details and API
// documentation in https://chromedevtools.github.io/devtools-protocol/tot,
// mirrored in https://github.com/ChromeDevTools/devtools-protocol).
package cdpgen

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const (
	cdpRoot   = "../../pkg/cdp"
	cdpURL    = "https://chromedevtools.github.io/devtools-protocol/tot"
	genURL    = "https://github.com/daabr/chrome-vision/cmd/cdpgen"
	importURL = "github.com/daabr/chrome-vision/pkg/cdp"
)

// CDP protocol file (browser_protocol.json, js_protocol.json).
type Protocol struct {
	Version Version
	Domains []Domain
}

// CDP protocol versions.
type Version struct {
	Major, Minor string
}

// CDP domain (group of related types, commands and events within a protocl).
type Domain struct {
	Domain                   string
	Description              *string
	Deprecated, Experimental bool
	Dependencies             []string
	Types                    []Type
	Commands                 []Command
	Events                   []Event
}

// CDP data type.
type Type struct {
	ID, Type                 string
	Description              *string
	Deprecated, Experimental bool
	Enum                     []string   // Optional, if the type is "string".
	Properties               []Property // Only if the type is "object".
	// Only when the type is "array"; in which case it will contains 1 key:
	// "type" (native JSON data type) or "$ref" (one of our custom types).
	Items map[string]string
}

// Property in a CDP type, or parameter in a CDP command or event.
type Property struct {
	Name                               string
	Description                        *string
	Deprecated, Experimental, Optional bool
	Type                               *string // Built-in JSON data types.
	Ref                                *string `json:"$ref"` // One of our custom types.
	// Only when type is "array", and only 1 key ("type" or "$ref").
	Items map[string]string
}

// CDP command.
type Command struct {
	Name                     string
	Description              *string
	Deprecated, Experimental bool
	Redirect                 *string
	Parameters               []Property
	Returns                  []Property
}

// CDP event.
type Event struct {
	Name                     string
	Description              *string
	Deprecated, Experimental bool
	Parameters               []Property
}

// Generate converts the given JSON-based CDP protocol to Go source code
// (one Go package per CDP domain, up to 4 files per package).
func Generate(p Protocol) {
	for _, d := range p.Domains {
		pkg := strings.ToLower(d.Domain)
		writeFile(cdpRoot+"/"+pkg, "doc.go", generateDoc(d, p.Version))
		if len(d.Types) > 0 {
			writeFile(cdpRoot+"/"+pkg, "types.go", generateTypes(d))
		}
		if len(d.Commands) > 0 {
			writeFile(cdpRoot+"/"+pkg, "commands.go", generateCommands(d))
		}
		if len(d.Events) > 0 {
			writeFile(cdpRoot+"/"+pkg, "events.go", generateEvents(d))
		}
	}
}

func writeFile(dir, file, content string) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(filepath.Join(dir, file))
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.WriteString(content); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

// generateImports generates Go import statements for types, commands and events files.
// Not really needed, given that we run "goimports" later, but still useful.
func generateImports(b *strings.Builder, standrd, github []string) {
	n := len(standrd) + len(github)
	if n > 0 {
		fmt.Fprint(b, "\n")
		if n > 1 {
			fmt.Fprintln(b, "import (")
		}
	}
	sort.Strings(standrd)
	for _, i := range standrd {
		generateImport(b, n, i)
	}
	if len(standrd) > 0 && len(github) > 0 {
		fmt.Fprint(b, "\n")
	}
	sort.Strings(github)
	for _, i := range github {
		generateImport(b, n, importURL+"/"+strings.ToLower(i))
	}
	if n > 1 {
		fmt.Fprintln(b, ")")
	}
}

func generateImport(b *strings.Builder, n int, s string) {
	if n == 1 {
		fmt.Fprint(b, "import ")
	} else {
		fmt.Fprint(b, "\t")
	}
	fmt.Fprintf(b, "%q\n", s)
}

// adjust applies Go naming conventions to CDP types and struct field names,
// and resolves circular dependencies (https://crbug.com/1193242).
func adjust(s string) string {
	// "-" enables word capitalization by string.Title(), "_" doesn't.
	s = strings.Title(strings.ReplaceAll(s, "_", "-"))
	// See https://golang.org/doc/effective_go#mixed-caps.
	s = strings.ReplaceAll(s, "-", "")

	re := regexp.MustCompile(`(Id)$`)
	s = re.ReplaceAllLiteralString(s, "ID")
	s = strings.ReplaceAll(s, "Url", "URL")

	if strings.Count(s, ".") > 0 {
		// Convert CDP domain prefixes to Go package prefixes.
		re := regexp.MustCompile(`^(.*)\.`)
		s = re.ReplaceAllStringFunc(s, strings.ToLower)

		// Resolve circular dependencies (https://crbug.com/1193242).
		s = strings.ReplaceAll(s, "network.TimeSinceEpoch", "cdp.TimeSinceEpoch")
		s = strings.ReplaceAll(s, "page.FrameID", "cdp.FrameID")
		s = strings.ReplaceAll(s, "target.TargetID", "cdp.TargetID")
	}
	return s
}

// convertType converts the given built-in JSON type to a Go data type.
func convertType(t string, arrayTypes map[string]string) string {
	switch t {
	case "any", "object":
		return "json.RawMessage"
	case "boolean":
		return "bool"
	case "integer":
		return "int64"
	case "number":
		return "float64"
	case "string":
		return "string"
	case "array":
		if item, ok := arrayTypes["type"]; ok {
			return "[]" + convertType(item, nil)
		}
		return "[]" + adjust(arrayTypes["$ref"])
	default:
		return ""
	}
}
