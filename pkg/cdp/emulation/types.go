package emulation

// Screen orientation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#type-ScreenOrientation
type ScreenOrientation struct {
	// Orientation type.
	Type string
	// Orientation angle.
	Angle int64
}

// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#type-DisplayFeature
type DisplayFeature struct {
	// Orientation of a display feature in relation to screen
	Orientation string
	// The offset from the screen origin in either the x (for vertical
	// orientation) or y (for horizontal orientation) direction.
	Offset int64
	// A display feature may mask content such that it is not physically
	// displayed - this length along with the offset describes this area.
	// A display feature that only splits content will have a 0 mask_length.
	MaskLength int64
}

// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#type-MediaFeature
type MediaFeature struct {
	Name  string
	Value string
}

// advance: If the scheduler runs out of immediate work, the virtual time base may fast forward to
// allow the next delayed task (if any) to run; pause: The virtual time base may not advance;
// pauseIfNetworkFetchesPending: The virtual time base may not advance if there are any pending
// resource fetches.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#type-VirtualTimePolicy
//
// This CDP type is experimental.
type VirtualTimePolicy string

// VirtualTimePolicy valid values.
const (
	VirtualTimePolicyAdvance                      VirtualTimePolicy = "advance"
	VirtualTimePolicyPause                        VirtualTimePolicy = "pause"
	VirtualTimePolicyPauseIfNetworkFetchesPending VirtualTimePolicy = "pauseIfNetworkFetchesPending"
)

// Used to specify User Agent Cient Hints to emulate. See https://wicg.github.io/ua-client-hints
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#type-UserAgentBrandVersion
//
// This CDP type is experimental.
type UserAgentBrandVersion struct {
	Brand   string
	Version string
}

// Used to specify User Agent Cient Hints to emulate. See https://wicg.github.io/ua-client-hints
// Missing optional values will be filled in by the target with what it would normally use.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#type-UserAgentMetadata
//
// This CDP type is experimental.
type UserAgentMetadata struct {
	Brands          []UserAgentBrandVersion `json:"brands,omitempty"`
	FullVersion     string                  `json:"fullVersion,omitempty"`
	Platform        string
	PlatformVersion string
	Architecture    string
	Model           string
	Mobile          bool
}

// Enum of image types that can be disabled.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#type-DisabledImageType
//
// This CDP type is experimental.
type DisabledImageType string

// DisabledImageType valid values.
const (
	DisabledImageTypeAvif DisabledImageType = "avif"
	DisabledImageTypeJxl  DisabledImageType = "jxl"
	DisabledImageTypeWebp DisabledImageType = "webp"
)
