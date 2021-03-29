package browser

// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#type-BrowserContextID
//
// This CDP type is experimental.
type BrowserContextID string

// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#type-WindowID
//
// This CDP type is experimental.
type WindowID int64

// The state of the browser window.
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

// Browser window bounds information
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

// Definition of PermissionDescriptor defined in the Permissions API:
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

// Browser command ids used by executeBrowserCommand.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#type-BrowserCommandId
//
// This CDP type is experimental.
type BrowserCommandID string

// BrowserCommandID valid values.
const (
	BrowserCommandIDOpenTabSearch  BrowserCommandID = "openTabSearch"
	BrowserCommandIDCloseTabSearch BrowserCommandID = "closeTabSearch"
)

// Chrome histogram bucket.
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

// Chrome histogram.
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
