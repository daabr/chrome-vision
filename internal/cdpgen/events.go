package cdpgen

import (
	"fmt"
	"strings"
)

func generateEvents(d Domain) string {
	b := new(strings.Builder)
	fmt.Fprintf(b, "package %s\n", strings.ToLower(d.Domain))

	for _, e := range d.Events {
		t := Type{
			ID:           e.Name,
			Type:         "object",
			Description:  e.Description,
			Deprecated:   e.Deprecated,
			Experimental: e.Experimental,
			Properties:   e.Parameters,
		}
		generateType(b, t, d.Domain, "event", nil)
	}
	return b.String()
}
