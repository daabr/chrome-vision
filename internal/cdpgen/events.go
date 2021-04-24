package cdpgen

import (
	"fmt"
	"strings"
)

func generateEvents(d Domain) string {
	b := new(strings.Builder)
	fmt.Fprintf(b, "package %s\n", strings.ToLower(d.Domain))

	for _, e := range d.Events {
		description := adjust(e.Name) + " asynchronous event."
		if e.Description != nil {
			description += "\n\n" + *e.Description
		}
		t := Type{
			ID:           e.Name,
			Type:         "object",
			Description:  &description,
			Deprecated:   e.Deprecated,
			Experimental: e.Experimental,
			Properties:   e.Parameters,
		}
		generateType(b, t, d.Domain, "event", nil)
	}
	return b.String()
}
