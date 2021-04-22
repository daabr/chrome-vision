// Package cdpgen generates the source code of the Go bindings for all the
// commands, events and types in the Chrome DevTools Protocol (CDP) - from the
// latest Chromium "tip-of-tree" (tot) definitions (see the details and API
// documentation in https://chromedevtools.github.io/devtools-protocol/tot,
// mirrored in https://github.com/ChromeDevTools/devtools-protocol).
package cdpgen

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	cdpRoot   = "../../pkg/cdp"
	cdpURL    = "https://chromedevtools.github.io/devtools-protocol/tot"
	genURL    = "https://github.com/daabr/chrome-vision/cmd/cdpgen"
	importURL = "github.com/daabr/chrome-vision/pkg/cdp"
)

// Protocol represents a CDP file (browser_protocol.json, js_protocol.json).
type Protocol struct {
	Version Version
	Domains []Domain
}

// Version represents a CDP protocol versions.
type Version struct {
	Major, Minor string
}

// Domain represents a CDP domain (group of related types, commands and
// events within a protocl).
type Domain struct {
	Domain                   string
	Description              *string
	Deprecated, Experimental bool
	Dependencies             []string // Unused, we rely on goimports.
	Types                    []Type
	Commands                 []Command
	Events                   []Event
}

// Type represents a CDP data type.
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

// Command in a CDP domain.
type Command struct {
	Name                     string
	Description              *string
	Deprecated, Experimental bool
	Redirect                 *string
	Parameters               []Property
	Returns                  []Property
}

// Event in a CDP domain.
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
		generateTypes(d) // Preparation for de-aliasing of built-in data types.
	}
	for _, d := range p.Domains {
		pkg := strings.ToLower(d.Domain)
		writeFile(pkg, "doc.go", generateDoc(d, p.Version))
		if len(d.Types) > 0 {
			writeFile(pkg, "types.go", generateTypes(d))
		}
		if len(d.Commands) > 0 {
			writeFile(pkg, "commands.go", generateCommands(d))
		}
		if len(d.Events) > 0 {
			writeFile(pkg, "events.go", generateEvents(d))
		}
	}
}

func writeFile(dir, file, content string) {
	if err := os.MkdirAll(filepath.Join(cdpRoot, dir), 0755); err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(filepath.Join(cdpRoot, dir, file))
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

// Apply Go naming conventions to CDP types and struct field names,
// and avoid circular dependencies (https://crbug.com/1193242).
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

		// Avoid circular dependencies (https://crbug.com/1193242).
		if t, ok := aliases[s]; ok {
			s = t
		}
	}
	return s
}

// Convert a JSON data type to a Go data type.
func transformType(t string, arrayTypes map[string]string) string {
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
			return "[]" + transformType(item, nil)
		}
		return "[]" + adjust(arrayTypes["$ref"])
	default:
		return ""
	}
}
