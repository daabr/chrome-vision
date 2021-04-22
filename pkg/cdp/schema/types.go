package schema

// Description of the protocol domain.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Schema/#type-Domain
type Domain struct {
	// Domain name.
	Name string
	// Domain version.
	Version string
}
