package schema

// Domain data type. Description of the protocol domain.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Schema/#type-Domain
type Domain struct {
	// Domain name.
	Name string `json:"name"`
	// Domain version.
	Version string `json:"version"`
}
