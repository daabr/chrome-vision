package cdpgen

import (
	"fmt"
	"strings"
)

// For de-aliasing of built-in data types.
var aliases = make(map[string]string)

func generateTypes(d Domain) string {
	b := new(strings.Builder)
	fmt.Fprintf(b, "package %s\n", strings.ToLower(d.Domain))

	for _, t := range d.Types {
		if t.Type != "array" && t.Type != "object" {
			aliases[adjust(t.ID)] = transformType(t.Type, nil)
			key := strings.ToLower(d.Domain) + "." + adjust(t.ID)
			aliases[key] = transformType(t.Type, nil)
		}
		description := adjust(t.ID) + " data type."
		if t.Description != nil {
			if strings.HasPrefix(*t.Description, t.ID) {
				description = *t.Description
			} else {
				description += " " + *t.Description
			}
		}
		t.Description = &description
		generateType(b, t, d.Domain, "type", nil)
	}
	return b.String()
}

// Generate structs - mainly for CDP types, but this is also reused for CDP
// commands and events. The redirect argument is relevant only for CDP commands.
func generateType(b *strings.Builder, t Type, domain, usage string, redirect *string) {
	// Comment.
	if t.Description != nil {
		fmt.Fprintf(b, "\n// %s\n", strings.ReplaceAll(*t.Description, "\n", "\n// "))
		fmt.Fprint(b, "//")
	}
	if usage != "return" {
		fmt.Fprintf(b, "\n// %s/%s/#%s-%s\n", cdpURL, domain, usage, t.ID)
	}
	if t.Deprecated || t.Experimental || redirect != nil {
		fmt.Fprintln(b, "//")
	}
	if t.Deprecated {
		fmt.Fprintf(b, "// This CDP %s is deprecated.\n", usage)
	}
	if t.Experimental {
		fmt.Fprintf(b, "// This CDP %s is experimental.\n", usage)
	}
	if redirect != nil {
		fmt.Fprintf(b, "// Use the `%s` domain instead.\n", *redirect)
	}

	// Type declaration.
	id := adjust(t.ID)
	switch t.Type {
	case "object":
		fmt.Fprintf(b, "type %s struct", id)
		if len(t.Properties) == 0 {
			fmt.Fprint(b, "{")
		} else {
			fmt.Fprint(b, " {\n")
		}
		for _, p := range t.Properties {
			pUsage := "property"
			if usage != "type" {
				pUsage = "parameter"
			}
			generateProperty(b, p, pUsage, domain)
		}
		fmt.Fprintln(b, "}")
	case "string":
		fmt.Fprintf(b, "type %s string\n", id)
		if len(t.Enum) > 0 {
			fmt.Fprintf(b, "\n// %s valid values.\n", id)
			fmt.Fprintln(b, "const (")
			for _, s := range t.Enum {
				fmt.Fprintf(b, "\t%s%s %s = %q\n", id, adjust(s), id, s)
			}
			fmt.Fprintln(b, ")")
		}
	default:
		fmt.Fprintf(b, "type %s %s\n", id, transformType(t.Type, t.Items))
	}
}

func generateProperty(b *strings.Builder, p Property, usage, domain string) {
	// Optional comment.
	if p.Description != nil {
		fmt.Fprintf(b, "\t// %s\n", strings.ReplaceAll(*p.Description, "\n", "\n\t// "))
		if p.Deprecated || p.Experimental {
			fmt.Fprintln(b, "\t//")
		}
	}
	if p.Deprecated {
		fmt.Fprintf(b, "\t// This CDP %s is deprecated.\n", usage)
	}
	if p.Experimental {
		fmt.Fprintf(b, "\t// This CDP %s is experimental.\n", usage)
	}

	// Struct field name.
	fmt.Fprintf(b, "\t%s ", adjust(p.Name))

	// Struct field type.
	if p.Type != nil {
		fmt.Fprint(b, transformType(*p.Type, p.Items))
	} else {
		r := strings.ReplaceAll(adjust(*p.Ref), strings.ToLower(domain)+".", "")
		if t, ok := aliases[r]; ok {
			r = t // De-alias built-in data types (https://crbug.com/1193242).
		}
		if p.Optional && r != "int64" && r != "float64" && r != "string" {
			fmt.Fprint(b, "*")
		}
		fmt.Fprint(b, r)
	}

	// JSON key hint.
	fmt.Fprintf(b, " `json:\"%s", p.Name)
	if p.Optional {
		fmt.Fprint(b, ",omitempty")
	}
	fmt.Fprint(b, "\"`\n")
}
