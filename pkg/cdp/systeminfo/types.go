package systeminfo

import "encoding/json"

// Describes a single graphics processor (GPU).
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#type-GPUDevice
type GPUDevice struct {
	// PCI ID of the GPU vendor, if available; 0 otherwise.
	VendorID float64
	// PCI ID of the GPU device, if available; 0 otherwise.
	DeviceID float64
	// Sub sys ID of the GPU, only available on Windows.
	SubSysID float64 `json:"subSysId,omitempty"`
	// Revision of the GPU, only available on Windows.
	Revision float64 `json:"revision,omitempty"`
	// String description of the GPU vendor, if the PCI ID is not available.
	VendorString string
	// String description of the GPU device, if the PCI ID is not available.
	DeviceString string
	// String description of the GPU driver vendor.
	DriverVendor string
	// String description of the GPU driver version.
	DriverVersion string
}

// Describes the width and height dimensions of an entity.
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#type-Size
type Size struct {
	// Width in pixels.
	Width int64
	// Height in pixels.
	Height int64
}

// Describes a supported video decoding profile with its associated minimum and
// maximum resolutions.
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#type-VideoDecodeAcceleratorCapability
type VideoDecodeAcceleratorCapability struct {
	// Video codec profile that is supported, e.g. VP9 Profile 2.
	Profile string
	// Maximum video dimensions in pixels supported for this |profile|.
	MaxResolution Size
	// Minimum video dimensions in pixels supported for this |profile|.
	MinResolution Size
}

// Describes a supported video encoding profile with its associated maximum
// resolution and maximum framerate.
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#type-VideoEncodeAcceleratorCapability
type VideoEncodeAcceleratorCapability struct {
	// Video codec profile that is supported, e.g H264 Main.
	Profile string
	// Maximum video dimensions in pixels supported for this |profile|.
	MaxResolution Size
	// Maximum encoding framerate in frames per second supported for this
	// |profile|, as fraction's numerator and denominator, e.g. 24/1 fps,
	// 24000/1001 fps, etc.
	MaxFramerateNumerator   int64
	MaxFramerateDenominator int64
}

// YUV subsampling type of the pixels of a given image.
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#type-SubsamplingFormat
type SubsamplingFormat string

// SubsamplingFormat valid values.
const (
	SubsamplingFormatYuv420 SubsamplingFormat = "yuv420"
	SubsamplingFormatYuv422 SubsamplingFormat = "yuv422"
	SubsamplingFormatYuv444 SubsamplingFormat = "yuv444"
)

// Image format of a given image.
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#type-ImageType
type ImageType string

// ImageType valid values.
const (
	ImageTypeJpeg    ImageType = "jpeg"
	ImageTypeWebp    ImageType = "webp"
	ImageTypeUnknown ImageType = "unknown"
)

// Describes a supported image decoding profile with its associated minimum and
// maximum resolutions and subsampling.
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#type-ImageDecodeAcceleratorCapability
type ImageDecodeAcceleratorCapability struct {
	// Image coded, e.g. Jpeg.
	ImageType string
	// Maximum supported dimensions of the image in pixels.
	MaxDimensions Size
	// Minimum supported dimensions of the image in pixels.
	MinDimensions Size
	// Optional array of supported subsampling formats, e.g. 4:2:0, if known.
	Subsamplings []SubsamplingFormat
}

// Provides information about the GPU(s) on the system.
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#type-GPUInfo
type GPUInfo struct {
	// The graphics devices on the system. Element 0 is the primary GPU.
	Devices []GPUDevice
	// An optional dictionary of additional GPU related attributes.
	AuxAttributes json.RawMessage `json:"auxAttributes,omitempty"`
	// An optional dictionary of graphics features and their status.
	FeatureStatus json.RawMessage `json:"featureStatus,omitempty"`
	// An optional array of GPU driver bug workarounds.
	DriverBugWorkarounds []string
	// Supported accelerated video decoding capabilities.
	VideoDecoding []VideoDecodeAcceleratorCapability
	// Supported accelerated video encoding capabilities.
	VideoEncoding []VideoEncodeAcceleratorCapability
	// Supported accelerated image decoding capabilities.
	ImageDecoding []ImageDecodeAcceleratorCapability
}

// Represents process info.
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#type-ProcessInfo
type ProcessInfo struct {
	// Specifies process type.
	Type string
	// Specifies process id.
	ID int64
	// Specifies cumulative CPU usage in seconds across all threads of the
	// process since the process start.
	CpuTime float64
}
