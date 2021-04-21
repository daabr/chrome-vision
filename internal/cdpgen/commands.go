package cdpgen

import (
	"fmt"
	"regexp"
	"strings"
)

func generateCommands(d Domain) string {
	b := new(strings.Builder)
	fmt.Fprintf(b, "package %s\n", strings.ToLower(d.Domain))
	generateImports(b, []string{"context", "encoding/json", "errors"}, d.Dependencies)

	for _, c := range d.Commands {
		// Don't ignore commands which are merely marked as deprecated,
		// but do ignore commands with an explicit replacement.
		if c.Redirect != nil {
			continue
		}

		// Receiver struct.
		cmd := adjust(c.Name)
		fmt.Fprintf(b, "\n// %s contains the parameters, and acts as\n", cmd)
		fmt.Fprintf(b, "// a Go receiver, for the CDP command `%s`.\n", c.Name)
		fmt.Fprint(b, "//")
		t := Type{
			ID:           c.Name,
			Type:         "object",
			Description:  c.Description,
			Deprecated:   c.Deprecated,
			Experimental: c.Experimental,
			Properties:   c.Parameters,
		}
		generateType(b, t, d.Domain, "method", c.Redirect)

		// Constructor function.
		fmt.Fprintf(b, "\n// New%s constructs a new %s struct instance, with\n", cmd, cmd)
		fmt.Fprintln(b, "// all (but only) the required parameters. Optional parameters")
		fmt.Fprintln(b, "// may be added using the builder-like methods below.")
		fmt.Fprintf(b, "//\n// %s/%s/#method-%s\n", cdpURL, d.Domain, c.Name)
		if c.Deprecated || c.Experimental || c.Redirect != nil {
			fmt.Fprintln(b, "//")
		}
		if c.Deprecated {
			fmt.Fprintln(b, "// This CDP method is deprecated.")
		}
		if c.Experimental {
			fmt.Fprintln(b, "// This CDP method is experimental.")
		}
		if c.Redirect != nil {
			fmt.Fprintf(b, "// Use the `%s` domain instead.\n", *c.Redirect)
		}
		fmt.Fprintf(b, "func New%s(", cmd)

		// Required parameters (in the constructor).
		var required, optional []Property
		for _, p := range c.Parameters {
			if p.Optional {
				optional = append(optional, p)
			} else {
				required = append(required, p)
			}
		}
		for i, p := range required {
			generateRequiredParameter(b, p)
			if i+1 < len(required) {
				fmt.Fprint(b, ", ")
			}
		}

		// Constructor method body.
		fmt.Fprintf(b, ") *%s {\n", cmd)
		fmt.Fprintf(b, "\treturn &%s{", cmd)
		if len(required) > 0 {
			fmt.Fprint(b, "\n")
			for _, p := range required {
				if p.Name == "range" || p.Name == "type" {
					// Don't use Go keywords as parameter names
					// (e.g. CSS.NewSetKeyframeKey, DOMDebugger.NewRemoveDOMBreakpoint).
					fmt.Fprintf(b, "\t\t%s: %s,\n", adjust(p.Name), string(p.Name[0]))
				} else {
					re := regexp.MustCompile(`(Id)$`)
					s := re.ReplaceAllLiteralString(p.Name, "ID")
					fmt.Fprintf(b, "\t\t%s: %s,\n", adjust(p.Name), s)
				}
			}
			fmt.Fprint(b, "\t")
		}
		fmt.Fprintln(b, "}\n}")

		// Optional parameters (as builder-like methods).
		for _, p := range optional {
			generateOptionalParameter(b, d.Domain, cmd, p)
		}

		// Return types.
		if len(c.Returns) > 0 {
			fmt.Fprintf(b, "\n// %sResponse contains the browser's response\n", cmd)
			fmt.Fprintf(b, "// to calling the %s CDP command with Do().\n", cmd)
			t := Type{
				ID:           cmd + "Response",
				Type:         "object",
				Description:  nil,
				Deprecated:   false,
				Experimental: false,
				Properties:   c.Returns,
			}
			generateType(b, t, d.Domain, "return", nil)
		}

		// Do() method to tie all of the above together.
		fmt.Fprintf(b, "\n// Do sends the %s CDP command to a browser,\n", cmd)
		fmt.Fprintln(b, "// and returns the browser's response.")
		fmt.Fprintf(b, "func (t *%s) Do(ctx context.Context) ", cmd)
		if len(c.Returns) == 0 {
			fmt.Fprintln(b, "error {")
		} else {
			fmt.Fprintf(b, "(*%sResponse, error) {\n", cmd)
		}
		if len(required)+len(optional) == 0 {
			fmt.Fprintf(b, "\tresponse, err := cdp.Send(ctx, \"%s.%s\", nil)\n", d.Domain, c.Name)
		} else {
			fmt.Fprintln(b, "\tb, err := json.Marshal(t)")
			fmt.Fprintln(b, "\tif err != nil {")
			if len(c.Returns) == 0 {
				fmt.Fprintln(b, "\t\treturn err")
			} else {
				fmt.Fprintln(b, "\t\treturn nil, err")
			}
			fmt.Fprintln(b, "\t}")
			fmt.Fprintf(b, "\tresponse, err := cdp.Send(ctx, \"%s.%s\", b)\n", d.Domain, c.Name)
		}
		fmt.Fprintln(b, "\tif err != nil {")
		if len(c.Returns) == 0 {
			fmt.Fprintln(b, "\t\treturn err")
		} else {
			fmt.Fprintln(b, "\t\treturn nil, err")
		}
		fmt.Fprintln(b, "\t}")
		fmt.Fprintln(b, "\tif response.Error != nil {")
		if len(c.Returns) == 0 {
			fmt.Fprintln(b, "\t\treturn errors.New(response.Error.Error())")
		} else {
			fmt.Fprintln(b, "\t\treturn nil, errors.New(response.Error.Error())")
		}
		fmt.Fprintln(b, "\t}")
		if len(c.Returns) == 0 {
			fmt.Fprintln(b, "\treturn nil")
		} else {
			fmt.Fprintf(b, "\tresult := &%sResponse{}\n", cmd)
			fmt.Fprintln(b, "\tif err := json.Unmarshal(response.Result, result); err != nil {")
			fmt.Fprintln(b, "\t\treturn nil, err")
			fmt.Fprintln(b, "\t}")
			fmt.Fprintln(b, "\treturn result, nil")
		}
		fmt.Fprintln(b, "}")
	}
	return b.String()
}

func generateRequiredParameter(b *strings.Builder, p Property) {
	// Name.
	if p.Name == "range" || p.Name == "type" {
		// Don't use Go keywords as a parameter names
		// (e.g. CSS.NewSetKeyframeKey, DOMDebugger.NewRemoveDOMBreakpoint).
		fmt.Fprintf(b, "%s ", string(p.Name[0]))
	} else {
		re := regexp.MustCompile(`(Id)$`)
		s := re.ReplaceAllLiteralString(p.Name, "ID")
		fmt.Fprintf(b, "%s ", s)
	}
	// Type.
	if p.Type != nil {
		fmt.Fprint(b, convertType(*p.Type, p.Items))
	} else {
		// De-alias built-in types (https://crbug.com/1193242).
		r := adjust(*p.Ref)
		if t, ok := aliases[r]; ok {
			r = t
		}
		if p.Optional {
			fmt.Fprintf(b, "*%s", r)
		} else {
			fmt.Fprint(b, r)
		}
	}
}

func generateOptionalParameter(b *strings.Builder, domain, cmd string, p Property) {
	// Comment.
	fmt.Fprintf(b, "\n// Set%s adds or modifies the value of the optional\n", adjust(p.Name))
	fmt.Fprintf(b, "// parameter `%s` in the %s CDP command.\n", p.Name, cmd)
	if p.Description != nil {
		fmt.Fprintf(b, "//\n// %s\n", strings.ReplaceAll(*p.Description, "\n", "\n// "))
	}
	if p.Deprecated || p.Experimental {
		fmt.Fprintln(b, "//")
	}
	if p.Deprecated {
		fmt.Fprintln(b, "// This CDP parameter is deprecated.")
	}
	if p.Experimental {
		fmt.Fprintln(b, "// This CDP parameter is experimental.")
	}

	// Method declaration.
	fmt.Fprintf(b, "func (t *%s) Set%s(v ", cmd, adjust(p.Name))
	if p.Type != nil {
		fmt.Fprint(b, convertType(*p.Type, p.Items))
	} else {
		r := strings.ReplaceAll(adjust(*p.Ref), strings.ToLower(domain)+".", "")
		// De-alias built-in types (https://crbug.com/1193242).
		if t, ok := aliases[r]; ok {
			r = t
		}
		fmt.Fprintf(b, "%s", r)
	}
	fmt.Fprintf(b, ") *%s {\n", cmd)

	// Method body.
	fmt.Fprintf(b, "\tt.%s = ", adjust(p.Name))
	if p.Type != nil {
		fmt.Fprintln(b, "v") // By value - built-in JSON types.
	} else {
		r := strings.ReplaceAll(adjust(*p.Ref), strings.ToLower(domain)+".", "")
		// De-alias built-in types (https://crbug.com/1193242).
		if t, ok := aliases[r]; ok {
			r = t
		}
		if r == "int64" || r == "float64" || r == "string" {
			fmt.Fprintln(b, "v") // By value - aliased built-in JSON types.
		} else {
			fmt.Fprintln(b, "&v") // By reference - CDP types.
		}
	}
	fmt.Fprintln(b, "\treturn t")
	fmt.Fprintln(b, "}")
}
