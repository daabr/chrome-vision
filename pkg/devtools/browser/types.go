package browser

// ContextID data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#type-BrowserContextID
//
// This CDP type is experimental.
type ContextID string

// WindowID data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#type-WindowID
//
// This CDP type is experimental.
type WindowID int64

// WindowState data type. The state of the browser window.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#type-WindowState
//
// This CDP type is experimental.
type WindowState string

// WindowState valid values.
const (
	WindowStateNormal     WindowState = "normal"
	WindowStateMinimized  WindowState = "minimized"
	WindowStateMaximized  WindowState = "maximized"
	WindowStateFullscreen WindowState = "fullscreen"
)

// String returns the WindowState value as a built-in string.
func (t WindowState) String() string {
	return string(t)
}

// Bounds data type. Browser window bounds information
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#type-Bounds
//
// This CDP type is experimental.
type Bounds struct {
	// The offset from the left edge of the screen to the window in pixels.
	Left int64 `json:"left,omitempty"`
	// The offset from the top edge of the screen to the window in pixels.
	Top int64 `json:"top,omitempty"`
	// The window width in pixels.
	Width int64 `json:"width,omitempty"`
	// The window height in pixels.
	Height int64 `json:"height,omitempty"`
	// The window state. Default to normal.
	WindowState *WindowState `json:"windowState,omitempty"`
}

// PermissionType data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#type-PermissionType
//
// This CDP type is experimental.
type PermissionType string

// PermissionType valid values.
const (
	PermissionTypeAccessibilityEvents      PermissionType = "accessibilityEvents"
	PermissionTypeAudioCapture             PermissionType = "audioCapture"
	PermissionTypeBackgroundSync           PermissionType = "backgroundSync"
	PermissionTypeBackgroundFetch          PermissionType = "backgroundFetch"
	PermissionTypeClipboardReadWrite       PermissionType = "clipboardReadWrite"
	PermissionTypeClipboardSanitizedWrite  PermissionType = "clipboardSanitizedWrite"
	PermissionTypeDisplayCapture           PermissionType = "displayCapture"
	PermissionTypeDurableStorage           PermissionType = "durableStorage"
	PermissionTypeFlash                    PermissionType = "flash"
	PermissionTypeGeolocation              PermissionType = "geolocation"
	PermissionTypeMidi                     PermissionType = "midi"
	PermissionTypeMidiSysex                PermissionType = "midiSysex"
	PermissionTypeNfc                      PermissionType = "nfc"
	PermissionTypeNotifications            PermissionType = "notifications"
	PermissionTypePaymentHandler           PermissionType = "paymentHandler"
	PermissionTypePeriodicBackgroundSync   PermissionType = "periodicBackgroundSync"
	PermissionTypeProtectedMediaIdentifier PermissionType = "protectedMediaIdentifier"
	PermissionTypeSensors                  PermissionType = "sensors"
	PermissionTypeVideoCapture             PermissionType = "videoCapture"
	PermissionTypeVideoCapturePanTiltZoom  PermissionType = "videoCapturePanTiltZoom"
	PermissionTypeIdleDetection            PermissionType = "idleDetection"
	PermissionTypeWakeLockScreen           PermissionType = "wakeLockScreen"
	PermissionTypeWakeLockSystem           PermissionType = "wakeLockSystem"
)

// String returns the PermissionType value as a built-in string.
func (t PermissionType) String() string {
	return string(t)
}

// PermissionSetting data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#type-PermissionSetting
//
// This CDP type is experimental.
type PermissionSetting string

// PermissionSetting valid values.
const (
	PermissionSettingGranted PermissionSetting = "granted"
	PermissionSettingDenied  PermissionSetting = "denied"
	PermissionSettingPrompt  PermissionSetting = "prompt"
)

// String returns the PermissionSetting value as a built-in string.
func (t PermissionSetting) String() string {
	return string(t)
}

// PermissionDescriptor data type. Definition of PermissionDescriptor defined in the Permissions API:
// https://w3c.github.io/permissions/#dictdef-permissiondescriptor.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#type-PermissionDescriptor
//
// This CDP type is experimental.
type PermissionDescriptor struct {
	// Name of permission.
	// See https://cs.chromium.org/chromium/src/third_party/blink/renderer/modules/permissions/permission_descriptor.idl for valid permission names.
	Name string `json:"name"`
	// For "midi" permission, may also specify sysex control.
	Sysex bool `json:"sysex,omitempty"`
	// For "push" permission, may specify userVisibleOnly.
	// Note that userVisibleOnly = true is the only currently supported type.
	UserVisibleOnly bool `json:"userVisibleOnly,omitempty"`
	// For "clipboard" permission, may specify allowWithoutSanitization.
	AllowWithoutSanitization bool `json:"allowWithoutSanitization,omitempty"`
	// For "camera" permission, may specify panTiltZoom.
	PanTiltZoom bool `json:"panTiltZoom,omitempty"`
}

// CommandID data type. Browser command ids used by executeBrowserCommand.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#type-BrowserCommandId
//
// This CDP type is experimental.
type CommandID string

// CommandID valid values.
const (
	CommandIDOpenTabSearch  CommandID = "openTabSearch"
	CommandIDCloseTabSearch CommandID = "closeTabSearch"
)

// String returns the CommandID value as a built-in string.
func (t CommandID) String() string {
	return string(t)
}

// Bucket data type. Chrome histogram bucket.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#type-Bucket
//
// This CDP type is experimental.
type Bucket struct {
	// Minimum value (inclusive).
	Low int64 `json:"low"`
	// Maximum value (exclusive).
	High int64 `json:"high"`
	// Number of samples.
	Count int64 `json:"count"`
}

// Histogram data type. Chrome histogram.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#type-Histogram
//
// This CDP type is experimental.
type Histogram struct {
	// Name.
	Name string `json:"name"`
	// Sum of sample values.
	Sum int64 `json:"sum"`
	// Total number of samples.
	Count int64 `json:"count"`
	// Buckets.
	Buckets []Bucket `json:"buckets"`
}
