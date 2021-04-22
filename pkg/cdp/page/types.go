package page

// Unique frame identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameId
type FrameID string

// Indicates whether a frame has been identified as an ad.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-AdFrameType
//
// This CDP type is experimental.
type AdFrameType string

// AdFrameType valid values.
const (
	AdFrameTypeNone  AdFrameType = "none"
	AdFrameTypeChild AdFrameType = "child"
	AdFrameTypeRoot  AdFrameType = "root"
)

// Indicates whether the frame is a secure context and why it is the case.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-SecureContextType
//
// This CDP type is experimental.
type SecureContextType string

// SecureContextType valid values.
const (
	SecureContextTypeSecure           SecureContextType = "Secure"
	SecureContextTypeSecureLocalhost  SecureContextType = "SecureLocalhost"
	SecureContextTypeInsecureScheme   SecureContextType = "InsecureScheme"
	SecureContextTypeInsecureAncestor SecureContextType = "InsecureAncestor"
)

// Indicates whether the frame is cross-origin isolated and why it is the case.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-CrossOriginIsolatedContextType
//
// This CDP type is experimental.
type CrossOriginIsolatedContextType string

// CrossOriginIsolatedContextType valid values.
const (
	CrossOriginIsolatedContextTypeIsolated                   CrossOriginIsolatedContextType = "Isolated"
	CrossOriginIsolatedContextTypeNotIsolated                CrossOriginIsolatedContextType = "NotIsolated"
	CrossOriginIsolatedContextTypeNotIsolatedFeatureDisabled CrossOriginIsolatedContextType = "NotIsolatedFeatureDisabled"
)

// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-GatedAPIFeatures
//
// This CDP type is experimental.
type GatedAPIFeatures string

// GatedAPIFeatures valid values.
const (
	GatedAPIFeaturesSharedArrayBuffers                GatedAPIFeatures = "SharedArrayBuffers"
	GatedAPIFeaturesSharedArrayBuffersTransferAllowed GatedAPIFeatures = "SharedArrayBuffersTransferAllowed"
	GatedAPIFeaturesPerformanceMeasureMemory          GatedAPIFeatures = "PerformanceMeasureMemory"
	GatedAPIFeaturesPerformanceProfile                GatedAPIFeatures = "PerformanceProfile"
)

// All Permissions Policy features. This enum should match the one defined
// in renderer/core/feature_policy/feature_policy_features.json5.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-PermissionsPolicyFeature
//
// This CDP type is experimental.
type PermissionsPolicyFeature string

// PermissionsPolicyFeature valid values.
const (
	PermissionsPolicyFeatureAccelerometer               PermissionsPolicyFeature = "accelerometer"
	PermissionsPolicyFeatureAmbientLightSensor          PermissionsPolicyFeature = "ambient-light-sensor"
	PermissionsPolicyFeatureAutoplay                    PermissionsPolicyFeature = "autoplay"
	PermissionsPolicyFeatureCamera                      PermissionsPolicyFeature = "camera"
	PermissionsPolicyFeatureChDpr                       PermissionsPolicyFeature = "ch-dpr"
	PermissionsPolicyFeatureChDeviceMemory              PermissionsPolicyFeature = "ch-device-memory"
	PermissionsPolicyFeatureChDownlink                  PermissionsPolicyFeature = "ch-downlink"
	PermissionsPolicyFeatureChEct                       PermissionsPolicyFeature = "ch-ect"
	PermissionsPolicyFeatureChLang                      PermissionsPolicyFeature = "ch-lang"
	PermissionsPolicyFeatureChRtt                       PermissionsPolicyFeature = "ch-rtt"
	PermissionsPolicyFeatureChUa                        PermissionsPolicyFeature = "ch-ua"
	PermissionsPolicyFeatureChUaArch                    PermissionsPolicyFeature = "ch-ua-arch"
	PermissionsPolicyFeatureChUaPlatform                PermissionsPolicyFeature = "ch-ua-platform"
	PermissionsPolicyFeatureChUaModel                   PermissionsPolicyFeature = "ch-ua-model"
	PermissionsPolicyFeatureChUaMobile                  PermissionsPolicyFeature = "ch-ua-mobile"
	PermissionsPolicyFeatureChUaFullVersion             PermissionsPolicyFeature = "ch-ua-full-version"
	PermissionsPolicyFeatureChUaPlatformVersion         PermissionsPolicyFeature = "ch-ua-platform-version"
	PermissionsPolicyFeatureChViewportWidth             PermissionsPolicyFeature = "ch-viewport-width"
	PermissionsPolicyFeatureChWidth                     PermissionsPolicyFeature = "ch-width"
	PermissionsPolicyFeatureClipboardRead               PermissionsPolicyFeature = "clipboard-read"
	PermissionsPolicyFeatureClipboardWrite              PermissionsPolicyFeature = "clipboard-write"
	PermissionsPolicyFeatureConversionMeasurement       PermissionsPolicyFeature = "conversion-measurement"
	PermissionsPolicyFeatureCrossOriginIsolated         PermissionsPolicyFeature = "cross-origin-isolated"
	PermissionsPolicyFeatureDisplayCapture              PermissionsPolicyFeature = "display-capture"
	PermissionsPolicyFeatureDocumentDomain              PermissionsPolicyFeature = "document-domain"
	PermissionsPolicyFeatureEncryptedMedia              PermissionsPolicyFeature = "encrypted-media"
	PermissionsPolicyFeatureExecutionWhileOutOfViewport PermissionsPolicyFeature = "execution-while-out-of-viewport"
	PermissionsPolicyFeatureExecutionWhileNotRendered   PermissionsPolicyFeature = "execution-while-not-rendered"
	PermissionsPolicyFeatureFocusWithoutUserActivation  PermissionsPolicyFeature = "focus-without-user-activation"
	PermissionsPolicyFeatureFullscreen                  PermissionsPolicyFeature = "fullscreen"
	PermissionsPolicyFeatureFrobulate                   PermissionsPolicyFeature = "frobulate"
	PermissionsPolicyFeatureGamepad                     PermissionsPolicyFeature = "gamepad"
	PermissionsPolicyFeatureGeolocation                 PermissionsPolicyFeature = "geolocation"
	PermissionsPolicyFeatureGyroscope                   PermissionsPolicyFeature = "gyroscope"
	PermissionsPolicyFeatureHid                         PermissionsPolicyFeature = "hid"
	PermissionsPolicyFeatureIdleDetection               PermissionsPolicyFeature = "idle-detection"
	PermissionsPolicyFeatureInterestCohort              PermissionsPolicyFeature = "interest-cohort"
	PermissionsPolicyFeatureMagnetometer                PermissionsPolicyFeature = "magnetometer"
	PermissionsPolicyFeatureMicrophone                  PermissionsPolicyFeature = "microphone"
	PermissionsPolicyFeatureMidi                        PermissionsPolicyFeature = "midi"
	PermissionsPolicyFeatureOtpCredentials              PermissionsPolicyFeature = "otp-credentials"
	PermissionsPolicyFeaturePayment                     PermissionsPolicyFeature = "payment"
	PermissionsPolicyFeaturePictureInPicture            PermissionsPolicyFeature = "picture-in-picture"
	PermissionsPolicyFeaturePublickeyCredentialsGet     PermissionsPolicyFeature = "publickey-credentials-get"
	PermissionsPolicyFeatureScreenWakeLock              PermissionsPolicyFeature = "screen-wake-lock"
	PermissionsPolicyFeatureSerial                      PermissionsPolicyFeature = "serial"
	PermissionsPolicyFeatureStorageAccessApi            PermissionsPolicyFeature = "storage-access-api"
	PermissionsPolicyFeatureSyncXhr                     PermissionsPolicyFeature = "sync-xhr"
	PermissionsPolicyFeatureTrustTokenRedemption        PermissionsPolicyFeature = "trust-token-redemption"
	PermissionsPolicyFeatureUsb                         PermissionsPolicyFeature = "usb"
	PermissionsPolicyFeatureVerticalScroll              PermissionsPolicyFeature = "vertical-scroll"
	PermissionsPolicyFeatureWebShare                    PermissionsPolicyFeature = "web-share"
	PermissionsPolicyFeatureXrSpatialTracking           PermissionsPolicyFeature = "xr-spatial-tracking"
)

// Reason for a permissions policy feature to be disabled.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-PermissionsPolicyBlockReason
//
// This CDP type is experimental.
type PermissionsPolicyBlockReason string

// PermissionsPolicyBlockReason valid values.
const (
	PermissionsPolicyBlockReasonHeader          PermissionsPolicyBlockReason = "Header"
	PermissionsPolicyBlockReasonIframeAttribute PermissionsPolicyBlockReason = "IframeAttribute"
)

// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-PermissionsPolicyBlockLocator
//
// This CDP type is experimental.
type PermissionsPolicyBlockLocator struct {
	FrameID     string
	BlockReason string
}

// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-PermissionsPolicyFeatureState
//
// This CDP type is experimental.
type PermissionsPolicyFeatureState struct {
	Feature string
	Allowed bool
	Locator *PermissionsPolicyBlockLocator `json:"locator,omitempty"`
}

// Information about the Frame on the page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-Frame
type Frame struct {
	// Frame unique identifier.
	ID string
	// Parent frame identifier.
	ParentID string `json:"parentId,omitempty"`
	// Identifier of the loader associated with this frame.
	LoaderID string
	// Frame's name as specified in the tag.
	Name string `json:"name,omitempty"`
	// Frame document's URL without fragment.
	URL string
	// Frame document's URL fragment including the '#'.
	//
	// This CDP property is experimental.
	URLFragment string `json:"urlFragment,omitempty"`
	// Frame document's registered domain, taking the public suffixes list into account.
	// Extracted from the Frame's url.
	// Example URLs: http://www.google.com/file.html -> "google.com"
	//               http://a.b.co.uk/file.html      -> "b.co.uk"
	//
	// This CDP property is experimental.
	DomainAndRegistry string
	// Frame document's security origin.
	SecurityOrigin string
	// Frame document's mimeType as determined by the browser.
	MimeType string
	// If the frame failed to load, this contains the URL that could not be loaded. Note that unlike url above, this URL may contain a fragment.
	//
	// This CDP property is experimental.
	UnreachableURL string `json:"unreachableUrl,omitempty"`
	// Indicates whether this frame was tagged as an ad.
	//
	// This CDP property is experimental.
	AdFrameType string `json:"adFrameType,omitempty"`
	// Indicates whether the main document is a secure context and explains why that is the case.
	//
	// This CDP property is experimental.
	SecureContextType string
	// Indicates whether this is a cross origin isolated context.
	//
	// This CDP property is experimental.
	CrossOriginIsolatedContextType string
	// Indicated which gated APIs / features are available.
	//
	// This CDP property is experimental.
	GatedAPIFeatures []GatedAPIFeatures
}

// Information about the Resource on the page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameResource
//
// This CDP type is experimental.
type FrameResource struct {
	// Resource URL.
	URL string
	// Type of this resource.
	Type string
	// Resource mimeType as determined by the browser.
	MimeType string
	// last-modified timestamp as reported by server.
	LastModified float64 `json:"lastModified,omitempty"`
	// Resource content size.
	ContentSize float64 `json:"contentSize,omitempty"`
	// True if the resource failed to load.
	Failed bool `json:"failed,omitempty"`
	// True if the resource was canceled during loading.
	Canceled bool `json:"canceled,omitempty"`
}

// Information about the Frame hierarchy along with their cached resources.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameResourceTree
//
// This CDP type is experimental.
type FrameResourceTree struct {
	// Frame information for this tree item.
	Frame Frame
	// Child frames.
	ChildFrames []FrameResourceTree `json:"childFrames,omitempty"`
	// Information about frame resources.
	Resources []FrameResource
}

// Information about the Frame hierarchy.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameTree
type FrameTree struct {
	// Frame information for this tree item.
	Frame Frame
	// Child frames.
	ChildFrames []FrameTree `json:"childFrames,omitempty"`
}

// Unique script identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-ScriptIdentifier
type ScriptIdentifier string

// Transition type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-TransitionType
type TransitionType string

// TransitionType valid values.
const (
	TransitionTypeLink             TransitionType = "link"
	TransitionTypeTyped            TransitionType = "typed"
	TransitionTypeAddressBar       TransitionType = "address_bar"
	TransitionTypeAutoBookmark     TransitionType = "auto_bookmark"
	TransitionTypeAutoSubframe     TransitionType = "auto_subframe"
	TransitionTypeManualSubframe   TransitionType = "manual_subframe"
	TransitionTypeGenerated        TransitionType = "generated"
	TransitionTypeAutoToplevel     TransitionType = "auto_toplevel"
	TransitionTypeFormSubmit       TransitionType = "form_submit"
	TransitionTypeReload           TransitionType = "reload"
	TransitionTypeKeyword          TransitionType = "keyword"
	TransitionTypeKeywordGenerated TransitionType = "keyword_generated"
	TransitionTypeOther            TransitionType = "other"
)

// Navigation history entry.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-NavigationEntry
type NavigationEntry struct {
	// Unique id of the navigation history entry.
	ID int64
	// URL of the navigation history entry.
	URL string
	// URL that the user typed in the url bar.
	UserTypedURL string
	// Title of the navigation history entry.
	Title string
	// Transition type.
	TransitionType string
}

// Screencast frame metadata.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-ScreencastFrameMetadata
//
// This CDP type is experimental.
type ScreencastFrameMetadata struct {
	// Top offset in DIP.
	OffsetTop float64
	// Page scale factor.
	PageScaleFactor float64
	// Device screen width in DIP.
	DeviceWidth float64
	// Device screen height in DIP.
	DeviceHeight float64
	// Position of horizontal scroll in CSS pixels.
	ScrollOffsetX float64
	// Position of vertical scroll in CSS pixels.
	ScrollOffsetY float64
	// Frame swap timestamp.
	Timestamp float64 `json:"timestamp,omitempty"`
}

// Javascript dialog type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-DialogType
type DialogType string

// DialogType valid values.
const (
	DialogTypeAlert        DialogType = "alert"
	DialogTypeConfirm      DialogType = "confirm"
	DialogTypePrompt       DialogType = "prompt"
	DialogTypeBeforeunload DialogType = "beforeunload"
)

// Error while paring app manifest.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-AppManifestError
type AppManifestError struct {
	// Error message.
	Message string
	// If criticial, this is a non-recoverable parse error.
	Critical int64
	// Error line.
	Line int64
	// Error column.
	Column int64
}

// Parsed app manifest properties.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-AppManifestParsedProperties
//
// This CDP type is experimental.
type AppManifestParsedProperties struct {
	// Computed scope value
	Scope string
}

// Layout viewport position and dimensions.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-LayoutViewport
type LayoutViewport struct {
	// Horizontal offset relative to the document (CSS pixels).
	PageX int64
	// Vertical offset relative to the document (CSS pixels).
	PageY int64
	// Width (CSS pixels), excludes scrollbar if present.
	ClientWidth int64
	// Height (CSS pixels), excludes scrollbar if present.
	ClientHeight int64
}

// Visual viewport position, dimensions, and scale.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-VisualViewport
type VisualViewport struct {
	// Horizontal offset relative to the layout viewport (CSS pixels).
	OffsetX float64
	// Vertical offset relative to the layout viewport (CSS pixels).
	OffsetY float64
	// Horizontal offset relative to the document (CSS pixels).
	PageX float64
	// Vertical offset relative to the document (CSS pixels).
	PageY float64
	// Width (CSS pixels), excludes scrollbar if present.
	ClientWidth float64
	// Height (CSS pixels), excludes scrollbar if present.
	ClientHeight float64
	// Scale relative to the ideal viewport (size at width=device-width).
	Scale float64
	// Page zoom factor (CSS to device independent pixels ratio).
	Zoom float64 `json:"zoom,omitempty"`
}

// Viewport for capturing screenshot.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-Viewport
type Viewport struct {
	// X offset in device independent pixels (dip).
	X float64
	// Y offset in device independent pixels (dip).
	Y float64
	// Rectangle width in device independent pixels (dip).
	Width float64
	// Rectangle height in device independent pixels (dip).
	Height float64
	// Page scale factor.
	Scale float64
}

// Generic font families collection.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FontFamilies
//
// This CDP type is experimental.
type FontFamilies struct {
	// The standard font-family.
	Standard string `json:"standard,omitempty"`
	// The fixed font-family.
	Fixed string `json:"fixed,omitempty"`
	// The serif font-family.
	Serif string `json:"serif,omitempty"`
	// The sansSerif font-family.
	SansSerif string `json:"sansSerif,omitempty"`
	// The cursive font-family.
	Cursive string `json:"cursive,omitempty"`
	// The fantasy font-family.
	Fantasy string `json:"fantasy,omitempty"`
	// The pictograph font-family.
	Pictograph string `json:"pictograph,omitempty"`
}

// Default font sizes.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FontSizes
//
// This CDP type is experimental.
type FontSizes struct {
	// Default standard font size.
	Standard int64 `json:"standard,omitempty"`
	// Default fixed font size.
	Fixed int64 `json:"fixed,omitempty"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-ClientNavigationReason
//
// This CDP type is experimental.
type ClientNavigationReason string

// ClientNavigationReason valid values.
const (
	ClientNavigationReasonFormSubmissionGet     ClientNavigationReason = "formSubmissionGet"
	ClientNavigationReasonFormSubmissionPost    ClientNavigationReason = "formSubmissionPost"
	ClientNavigationReasonHttpHeaderRefresh     ClientNavigationReason = "httpHeaderRefresh"
	ClientNavigationReasonScriptInitiated       ClientNavigationReason = "scriptInitiated"
	ClientNavigationReasonMetaTagRefresh        ClientNavigationReason = "metaTagRefresh"
	ClientNavigationReasonPageBlockInterstitial ClientNavigationReason = "pageBlockInterstitial"
	ClientNavigationReasonReload                ClientNavigationReason = "reload"
	ClientNavigationReasonAnchorClick           ClientNavigationReason = "anchorClick"
)

// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-ClientNavigationDisposition
//
// This CDP type is experimental.
type ClientNavigationDisposition string

// ClientNavigationDisposition valid values.
const (
	ClientNavigationDispositionCurrentTab ClientNavigationDisposition = "currentTab"
	ClientNavigationDispositionNewTab     ClientNavigationDisposition = "newTab"
	ClientNavigationDispositionNewWindow  ClientNavigationDisposition = "newWindow"
	ClientNavigationDispositionDownload   ClientNavigationDisposition = "download"
)

// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-InstallabilityErrorArgument
//
// This CDP type is experimental.
type InstallabilityErrorArgument struct {
	// Argument name (e.g. name:'minimum-icon-size-in-pixels').
	Name string
	// Argument value (e.g. value:'64').
	Value string
}

// The installability error
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-InstallabilityError
//
// This CDP type is experimental.
type InstallabilityError struct {
	// The error id (e.g. 'manifest-missing-suitable-icon').
	ErrorID string
	// The list of error arguments (e.g. {name:'minimum-icon-size-in-pixels', value:'64'}).
	ErrorArguments []InstallabilityErrorArgument
}

// The referring-policy used for the navigation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-ReferrerPolicy
//
// This CDP type is experimental.
type ReferrerPolicy string

// ReferrerPolicy valid values.
const (
	ReferrerPolicyNoReferrer                  ReferrerPolicy = "noReferrer"
	ReferrerPolicyNoReferrerWhenDowngrade     ReferrerPolicy = "noReferrerWhenDowngrade"
	ReferrerPolicyOrigin                      ReferrerPolicy = "origin"
	ReferrerPolicyOriginWhenCrossOrigin       ReferrerPolicy = "originWhenCrossOrigin"
	ReferrerPolicySameOrigin                  ReferrerPolicy = "sameOrigin"
	ReferrerPolicyStrictOrigin                ReferrerPolicy = "strictOrigin"
	ReferrerPolicyStrictOriginWhenCrossOrigin ReferrerPolicy = "strictOriginWhenCrossOrigin"
	ReferrerPolicyUnsafeURL                   ReferrerPolicy = "unsafeUrl"
)

// Per-script compilation cache parameters for `Page.produceCompilationCache`
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-CompilationCacheParams
//
// This CDP type is experimental.
type CompilationCacheParams struct {
	// The URL of the script to produce a compilation cache entry for.
	URL string
	// A hint to the backend whether eager compilation is recommended.
	// (the actual compilation mode used is upon backend discretion).
	Eager bool `json:"eager,omitempty"`
}

// The type of a frameNavigated event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-NavigationType
//
// This CDP type is experimental.
type NavigationType string

// NavigationType valid values.
const (
	NavigationTypeNavigation              NavigationType = "Navigation"
	NavigationTypeBackForwardCacheRestore NavigationType = "BackForwardCacheRestore"
)
