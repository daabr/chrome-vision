package page

import "github.com/daabr/chrome-vision/pkg/devtools/network"

// FrameID data type. Unique frame identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameId
type FrameID string

// AdFrameType data type. Indicates whether a frame has been identified as an ad.
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

// String returns the AdFrameType value as a built-in string.
func (t AdFrameType) String() string {
	return string(t)
}

// AdFrameExplanation data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-AdFrameExplanation
//
// This CDP type is experimental.
type AdFrameExplanation string

// AdFrameExplanation valid values.
const (
	AdFrameExplanationParentIsAd          AdFrameExplanation = "ParentIsAd"
	AdFrameExplanationCreatedByAdScript   AdFrameExplanation = "CreatedByAdScript"
	AdFrameExplanationMatchedBlockingRule AdFrameExplanation = "MatchedBlockingRule"
)

// String returns the AdFrameExplanation value as a built-in string.
func (t AdFrameExplanation) String() string {
	return string(t)
}

// AdFrameStatus data type. Indicates whether a frame has been identified as an ad and why.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-AdFrameStatus
//
// This CDP type is experimental.
type AdFrameStatus struct {
	AdFrameType  AdFrameType          `json:"adFrameType"`
	Explanations []AdFrameExplanation `json:"explanations,omitempty"`
}

// SecureContextType data type. Indicates whether the frame is a secure context and why it is the case.
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

// String returns the SecureContextType value as a built-in string.
func (t SecureContextType) String() string {
	return string(t)
}

// CrossOriginIsolatedContextType data type. Indicates whether the frame is cross-origin isolated and why it is the case.
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

// String returns the CrossOriginIsolatedContextType value as a built-in string.
func (t CrossOriginIsolatedContextType) String() string {
	return string(t)
}

// GatedAPIFeatures data type.
//
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

// String returns the GatedAPIFeatures value as a built-in string.
func (t GatedAPIFeatures) String() string {
	return string(t)
}

// PermissionsPolicyFeature data type. All Permissions Policy features. This enum should match the one defined
// in third_party/blink/renderer/core/permissions_policy/permissions_policy_features.json5.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-PermissionsPolicyFeature
//
// This CDP type is experimental.
type PermissionsPolicyFeature string

// PermissionsPolicyFeature valid values.
const (
	PermissionsPolicyFeatureAccelerometer               PermissionsPolicyFeature = "accelerometer"
	PermissionsPolicyFeatureAmbientLightSensor          PermissionsPolicyFeature = "ambient-light-sensor"
	PermissionsPolicyFeatureAttributionReporting        PermissionsPolicyFeature = "attribution-reporting"
	PermissionsPolicyFeatureAutoplay                    PermissionsPolicyFeature = "autoplay"
	PermissionsPolicyFeatureCamera                      PermissionsPolicyFeature = "camera"
	PermissionsPolicyFeatureChDpr                       PermissionsPolicyFeature = "ch-dpr"
	PermissionsPolicyFeatureChDeviceMemory              PermissionsPolicyFeature = "ch-device-memory"
	PermissionsPolicyFeatureChDownlink                  PermissionsPolicyFeature = "ch-downlink"
	PermissionsPolicyFeatureChEct                       PermissionsPolicyFeature = "ch-ect"
	PermissionsPolicyFeatureChLang                      PermissionsPolicyFeature = "ch-lang"
	PermissionsPolicyFeatureChPrefersColorScheme        PermissionsPolicyFeature = "ch-prefers-color-scheme"
	PermissionsPolicyFeatureChRtt                       PermissionsPolicyFeature = "ch-rtt"
	PermissionsPolicyFeatureChUa                        PermissionsPolicyFeature = "ch-ua"
	PermissionsPolicyFeatureChUaArch                    PermissionsPolicyFeature = "ch-ua-arch"
	PermissionsPolicyFeatureChUaBitness                 PermissionsPolicyFeature = "ch-ua-bitness"
	PermissionsPolicyFeatureChUaPlatform                PermissionsPolicyFeature = "ch-ua-platform"
	PermissionsPolicyFeatureChUaModel                   PermissionsPolicyFeature = "ch-ua-model"
	PermissionsPolicyFeatureChUaMobile                  PermissionsPolicyFeature = "ch-ua-mobile"
	PermissionsPolicyFeatureChUaFullVersion             PermissionsPolicyFeature = "ch-ua-full-version"
	PermissionsPolicyFeatureChUaPlatformVersion         PermissionsPolicyFeature = "ch-ua-platform-version"
	PermissionsPolicyFeatureChUaReduced                 PermissionsPolicyFeature = "ch-ua-reduced"
	PermissionsPolicyFeatureChViewportHeight            PermissionsPolicyFeature = "ch-viewport-height"
	PermissionsPolicyFeatureChViewportWidth             PermissionsPolicyFeature = "ch-viewport-width"
	PermissionsPolicyFeatureChWidth                     PermissionsPolicyFeature = "ch-width"
	PermissionsPolicyFeatureClipboardRead               PermissionsPolicyFeature = "clipboard-read"
	PermissionsPolicyFeatureClipboardWrite              PermissionsPolicyFeature = "clipboard-write"
	PermissionsPolicyFeatureCrossOriginIsolated         PermissionsPolicyFeature = "cross-origin-isolated"
	PermissionsPolicyFeatureDirectSockets               PermissionsPolicyFeature = "direct-sockets"
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
	PermissionsPolicyFeatureSharedAutofill              PermissionsPolicyFeature = "shared-autofill"
	PermissionsPolicyFeatureStorageAccessAPI            PermissionsPolicyFeature = "storage-access-api"
	PermissionsPolicyFeatureSyncXhr                     PermissionsPolicyFeature = "sync-xhr"
	PermissionsPolicyFeatureTrustTokenRedemption        PermissionsPolicyFeature = "trust-token-redemption"
	PermissionsPolicyFeatureUsb                         PermissionsPolicyFeature = "usb"
	PermissionsPolicyFeatureVerticalScroll              PermissionsPolicyFeature = "vertical-scroll"
	PermissionsPolicyFeatureWebShare                    PermissionsPolicyFeature = "web-share"
	PermissionsPolicyFeatureWindowPlacement             PermissionsPolicyFeature = "window-placement"
	PermissionsPolicyFeatureXrSpatialTracking           PermissionsPolicyFeature = "xr-spatial-tracking"
)

// String returns the PermissionsPolicyFeature value as a built-in string.
func (t PermissionsPolicyFeature) String() string {
	return string(t)
}

// PermissionsPolicyBlockReason data type. Reason for a permissions policy feature to be disabled.
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

// String returns the PermissionsPolicyBlockReason value as a built-in string.
func (t PermissionsPolicyBlockReason) String() string {
	return string(t)
}

// PermissionsPolicyBlockLocator data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-PermissionsPolicyBlockLocator
//
// This CDP type is experimental.
type PermissionsPolicyBlockLocator struct {
	FrameID     string                       `json:"frameId"`
	BlockReason PermissionsPolicyBlockReason `json:"blockReason"`
}

// PermissionsPolicyFeatureState data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-PermissionsPolicyFeatureState
//
// This CDP type is experimental.
type PermissionsPolicyFeatureState struct {
	Feature PermissionsPolicyFeature       `json:"feature"`
	Allowed bool                           `json:"allowed"`
	Locator *PermissionsPolicyBlockLocator `json:"locator,omitempty"`
}

// OriginTrialTokenStatus data type. Origin Trial(https://www.chromium.org/blink/origin-trials) support.
// Status for an Origin Trial token.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-OriginTrialTokenStatus
//
// This CDP type is experimental.
type OriginTrialTokenStatus string

// OriginTrialTokenStatus valid values.
const (
	OriginTrialTokenStatusSuccess                OriginTrialTokenStatus = "Success"
	OriginTrialTokenStatusNotSupported           OriginTrialTokenStatus = "NotSupported"
	OriginTrialTokenStatusInsecure               OriginTrialTokenStatus = "Insecure"
	OriginTrialTokenStatusExpired                OriginTrialTokenStatus = "Expired"
	OriginTrialTokenStatusWrongOrigin            OriginTrialTokenStatus = "WrongOrigin"
	OriginTrialTokenStatusInvalidSignature       OriginTrialTokenStatus = "InvalidSignature"
	OriginTrialTokenStatusMalformed              OriginTrialTokenStatus = "Malformed"
	OriginTrialTokenStatusWrongVersion           OriginTrialTokenStatus = "WrongVersion"
	OriginTrialTokenStatusFeatureDisabled        OriginTrialTokenStatus = "FeatureDisabled"
	OriginTrialTokenStatusTokenDisabled          OriginTrialTokenStatus = "TokenDisabled"
	OriginTrialTokenStatusFeatureDisabledForUser OriginTrialTokenStatus = "FeatureDisabledForUser"
)

// String returns the OriginTrialTokenStatus value as a built-in string.
func (t OriginTrialTokenStatus) String() string {
	return string(t)
}

// OriginTrialStatus data type. Status for an Origin Trial.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-OriginTrialStatus
//
// This CDP type is experimental.
type OriginTrialStatus string

// OriginTrialStatus valid values.
const (
	OriginTrialStatusEnabled               OriginTrialStatus = "Enabled"
	OriginTrialStatusValidTokenNotProvided OriginTrialStatus = "ValidTokenNotProvided"
	OriginTrialStatusOSNotSupported        OriginTrialStatus = "OSNotSupported"
	OriginTrialStatusTrialNotAllowed       OriginTrialStatus = "TrialNotAllowed"
)

// String returns the OriginTrialStatus value as a built-in string.
func (t OriginTrialStatus) String() string {
	return string(t)
}

// OriginTrialUsageRestriction data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-OriginTrialUsageRestriction
//
// This CDP type is experimental.
type OriginTrialUsageRestriction string

// OriginTrialUsageRestriction valid values.
const (
	OriginTrialUsageRestrictionNone   OriginTrialUsageRestriction = "None"
	OriginTrialUsageRestrictionSubset OriginTrialUsageRestriction = "Subset"
)

// String returns the OriginTrialUsageRestriction value as a built-in string.
func (t OriginTrialUsageRestriction) String() string {
	return string(t)
}

// OriginTrialToken data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-OriginTrialToken
//
// This CDP type is experimental.
type OriginTrialToken struct {
	Origin           string                      `json:"origin"`
	MatchSubDomains  bool                        `json:"matchSubDomains"`
	TrialName        string                      `json:"trialName"`
	ExpiryTime       float64                     `json:"expiryTime"`
	IsThirdParty     bool                        `json:"isThirdParty"`
	UsageRestriction OriginTrialUsageRestriction `json:"usageRestriction"`
}

// OriginTrialTokenWithStatus data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-OriginTrialTokenWithStatus
//
// This CDP type is experimental.
type OriginTrialTokenWithStatus struct {
	RawTokenText string `json:"rawTokenText"`
	// `parsedToken` is present only when the token is extractable and
	// parsable.
	ParsedToken *OriginTrialToken      `json:"parsedToken,omitempty"`
	Status      OriginTrialTokenStatus `json:"status"`
}

// OriginTrial data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-OriginTrial
//
// This CDP type is experimental.
type OriginTrial struct {
	TrialName        string                       `json:"trialName"`
	Status           OriginTrialStatus            `json:"status"`
	TokensWithStatus []OriginTrialTokenWithStatus `json:"tokensWithStatus"`
}

// Frame data type. Information about the Frame on the page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-Frame
type Frame struct {
	// Frame unique identifier.
	ID string `json:"id"`
	// Parent frame identifier.
	ParentID string `json:"parentId,omitempty"`
	// Identifier of the loader associated with this frame.
	LoaderID string `json:"loaderId"`
	// Frame's name as specified in the tag.
	Name string `json:"name,omitempty"`
	// Frame document's URL without fragment.
	URL string `json:"url"`
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
	DomainAndRegistry string `json:"domainAndRegistry"`
	// Frame document's security origin.
	SecurityOrigin string `json:"securityOrigin"`
	// Frame document's mimeType as determined by the browser.
	MimeType string `json:"mimeType"`
	// If the frame failed to load, this contains the URL that could not be loaded. Note that unlike url above, this URL may contain a fragment.
	//
	// This CDP property is experimental.
	UnreachableURL string `json:"unreachableUrl,omitempty"`
	// Indicates whether this frame was tagged as an ad and why.
	//
	// This CDP property is experimental.
	AdFrameStatus *AdFrameStatus `json:"adFrameStatus,omitempty"`
	// Indicates whether the main document is a secure context and explains why that is the case.
	//
	// This CDP property is experimental.
	SecureContextType SecureContextType `json:"secureContextType"`
	// Indicates whether this is a cross origin isolated context.
	//
	// This CDP property is experimental.
	CrossOriginIsolatedContextType CrossOriginIsolatedContextType `json:"crossOriginIsolatedContextType"`
	// Indicated which gated APIs / features are available.
	//
	// This CDP property is experimental.
	GatedAPIFeatures []GatedAPIFeatures `json:"gatedAPIFeatures"`
}

// FrameResource data type. Information about the Resource on the page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameResource
//
// This CDP type is experimental.
type FrameResource struct {
	// Resource URL.
	URL string `json:"url"`
	// Type of this resource.
	Type network.ResourceType `json:"type"`
	// Resource mimeType as determined by the browser.
	MimeType string `json:"mimeType"`
	// last-modified timestamp as reported by server.
	LastModified float64 `json:"lastModified,omitempty"`
	// Resource content size.
	ContentSize float64 `json:"contentSize,omitempty"`
	// True if the resource failed to load.
	Failed bool `json:"failed,omitempty"`
	// True if the resource was canceled during loading.
	Canceled bool `json:"canceled,omitempty"`
}

// FrameResourceTree data type. Information about the Frame hierarchy along with their cached resources.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameResourceTree
//
// This CDP type is experimental.
type FrameResourceTree struct {
	// Frame information for this tree item.
	Frame Frame `json:"frame"`
	// Child frames.
	ChildFrames []FrameResourceTree `json:"childFrames,omitempty"`
	// Information about frame resources.
	Resources []FrameResource `json:"resources"`
}

// FrameTree data type. Information about the Frame hierarchy.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameTree
type FrameTree struct {
	// Frame information for this tree item.
	Frame Frame `json:"frame"`
	// Child frames.
	ChildFrames []FrameTree `json:"childFrames,omitempty"`
}

// ScriptIdentifier data type. Unique script identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-ScriptIdentifier
type ScriptIdentifier string

// TransitionType data type. Transition type.
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

// String returns the TransitionType value as a built-in string.
func (t TransitionType) String() string {
	return string(t)
}

// NavigationEntry data type. Navigation history entry.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-NavigationEntry
type NavigationEntry struct {
	// Unique id of the navigation history entry.
	ID int64 `json:"id"`
	// URL of the navigation history entry.
	URL string `json:"url"`
	// URL that the user typed in the url bar.
	UserTypedURL string `json:"userTypedURL"`
	// Title of the navigation history entry.
	Title string `json:"title"`
	// Transition type.
	TransitionType TransitionType `json:"transitionType"`
}

// ScreencastFrameMetadata data type. Screencast frame metadata.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-ScreencastFrameMetadata
//
// This CDP type is experimental.
type ScreencastFrameMetadata struct {
	// Top offset in DIP.
	OffsetTop float64 `json:"offsetTop"`
	// Page scale factor.
	PageScaleFactor float64 `json:"pageScaleFactor"`
	// Device screen width in DIP.
	DeviceWidth float64 `json:"deviceWidth"`
	// Device screen height in DIP.
	DeviceHeight float64 `json:"deviceHeight"`
	// Position of horizontal scroll in CSS pixels.
	ScrollOffsetX float64 `json:"scrollOffsetX"`
	// Position of vertical scroll in CSS pixels.
	ScrollOffsetY float64 `json:"scrollOffsetY"`
	// Frame swap timestamp.
	Timestamp float64 `json:"timestamp,omitempty"`
}

// DialogType data type. Javascript dialog type.
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

// String returns the DialogType value as a built-in string.
func (t DialogType) String() string {
	return string(t)
}

// AppManifestError data type. Error while paring app manifest.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-AppManifestError
type AppManifestError struct {
	// Error message.
	Message string `json:"message"`
	// If criticial, this is a non-recoverable parse error.
	Critical int64 `json:"critical"`
	// Error line.
	Line int64 `json:"line"`
	// Error column.
	Column int64 `json:"column"`
}

// AppManifestParsedProperties data type. Parsed app manifest properties.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-AppManifestParsedProperties
//
// This CDP type is experimental.
type AppManifestParsedProperties struct {
	// Computed scope value
	Scope string `json:"scope"`
}

// LayoutViewport data type. Layout viewport position and dimensions.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-LayoutViewport
type LayoutViewport struct {
	// Horizontal offset relative to the document (CSS pixels).
	PageX int64 `json:"pageX"`
	// Vertical offset relative to the document (CSS pixels).
	PageY int64 `json:"pageY"`
	// Width (CSS pixels), excludes scrollbar if present.
	ClientWidth int64 `json:"clientWidth"`
	// Height (CSS pixels), excludes scrollbar if present.
	ClientHeight int64 `json:"clientHeight"`
}

// VisualViewport data type. Visual viewport position, dimensions, and scale.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-VisualViewport
type VisualViewport struct {
	// Horizontal offset relative to the layout viewport (CSS pixels).
	OffsetX float64 `json:"offsetX"`
	// Vertical offset relative to the layout viewport (CSS pixels).
	OffsetY float64 `json:"offsetY"`
	// Horizontal offset relative to the document (CSS pixels).
	PageX float64 `json:"pageX"`
	// Vertical offset relative to the document (CSS pixels).
	PageY float64 `json:"pageY"`
	// Width (CSS pixels), excludes scrollbar if present.
	ClientWidth float64 `json:"clientWidth"`
	// Height (CSS pixels), excludes scrollbar if present.
	ClientHeight float64 `json:"clientHeight"`
	// Scale relative to the ideal viewport (size at width=device-width).
	Scale float64 `json:"scale"`
	// Page zoom factor (CSS to device independent pixels ratio).
	Zoom float64 `json:"zoom,omitempty"`
}

// Viewport for capturing screenshot.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-Viewport
type Viewport struct {
	// X offset in device independent pixels (dip).
	X float64 `json:"x"`
	// Y offset in device independent pixels (dip).
	Y float64 `json:"y"`
	// Rectangle width in device independent pixels (dip).
	Width float64 `json:"width"`
	// Rectangle height in device independent pixels (dip).
	Height float64 `json:"height"`
	// Page scale factor.
	Scale float64 `json:"scale"`
}

// FontFamilies data type. Generic font families collection.
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

// FontSizes data type. Default font sizes.
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

// ClientNavigationReason data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-ClientNavigationReason
//
// This CDP type is experimental.
type ClientNavigationReason string

// ClientNavigationReason valid values.
const (
	ClientNavigationReasonFormSubmissionGet     ClientNavigationReason = "formSubmissionGet"
	ClientNavigationReasonFormSubmissionPost    ClientNavigationReason = "formSubmissionPost"
	ClientNavigationReasonHTTPHeaderRefresh     ClientNavigationReason = "httpHeaderRefresh"
	ClientNavigationReasonScriptInitiated       ClientNavigationReason = "scriptInitiated"
	ClientNavigationReasonMetaTagRefresh        ClientNavigationReason = "metaTagRefresh"
	ClientNavigationReasonPageBlockInterstitial ClientNavigationReason = "pageBlockInterstitial"
	ClientNavigationReasonReload                ClientNavigationReason = "reload"
	ClientNavigationReasonAnchorClick           ClientNavigationReason = "anchorClick"
)

// String returns the ClientNavigationReason value as a built-in string.
func (t ClientNavigationReason) String() string {
	return string(t)
}

// ClientNavigationDisposition data type.
//
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

// String returns the ClientNavigationDisposition value as a built-in string.
func (t ClientNavigationDisposition) String() string {
	return string(t)
}

// InstallabilityErrorArgument data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-InstallabilityErrorArgument
//
// This CDP type is experimental.
type InstallabilityErrorArgument struct {
	// Argument name (e.g. name:'minimum-icon-size-in-pixels').
	Name string `json:"name"`
	// Argument value (e.g. value:'64').
	Value string `json:"value"`
}

// InstallabilityError data type. The installability error
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-InstallabilityError
//
// This CDP type is experimental.
type InstallabilityError struct {
	// The error id (e.g. 'manifest-missing-suitable-icon').
	ErrorID string `json:"errorId"`
	// The list of error arguments (e.g. {name:'minimum-icon-size-in-pixels', value:'64'}).
	ErrorArguments []InstallabilityErrorArgument `json:"errorArguments"`
}

// ReferrerPolicy data type. The referring-policy used for the navigation.
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

// String returns the ReferrerPolicy value as a built-in string.
func (t ReferrerPolicy) String() string {
	return string(t)
}

// CompilationCacheParams data type. Per-script compilation cache parameters for `Page.produceCompilationCache`
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-CompilationCacheParams
//
// This CDP type is experimental.
type CompilationCacheParams struct {
	// The URL of the script to produce a compilation cache entry for.
	URL string `json:"url"`
	// A hint to the backend whether eager compilation is recommended.
	// (the actual compilation mode used is upon backend discretion).
	Eager bool `json:"eager,omitempty"`
}

// NavigationType data type. The type of a frameNavigated event.
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

// String returns the NavigationType value as a built-in string.
func (t NavigationType) String() string {
	return string(t)
}

// BackForwardCacheNotRestoredReason data type. List of not restored reasons for back-forward cache.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-BackForwardCacheNotRestoredReason
//
// This CDP type is experimental.
type BackForwardCacheNotRestoredReason string

// BackForwardCacheNotRestoredReason valid values.
const (
	BackForwardCacheNotRestoredReasonNotMainFrame                                             BackForwardCacheNotRestoredReason = "NotMainFrame"
	BackForwardCacheNotRestoredReasonBackForwardCacheDisabled                                 BackForwardCacheNotRestoredReason = "BackForwardCacheDisabled"
	BackForwardCacheNotRestoredReasonRelatedActiveContentsExist                               BackForwardCacheNotRestoredReason = "RelatedActiveContentsExist"
	BackForwardCacheNotRestoredReasonHTTPStatusNotOK                                          BackForwardCacheNotRestoredReason = "HTTPStatusNotOK"
	BackForwardCacheNotRestoredReasonSchemeNotHTTPOrHTTPS                                     BackForwardCacheNotRestoredReason = "SchemeNotHTTPOrHTTPS"
	BackForwardCacheNotRestoredReasonLoading                                                  BackForwardCacheNotRestoredReason = "Loading"
	BackForwardCacheNotRestoredReasonWasGrantedMediaAccess                                    BackForwardCacheNotRestoredReason = "WasGrantedMediaAccess"
	BackForwardCacheNotRestoredReasonDisableForRenderFrameHostCalled                          BackForwardCacheNotRestoredReason = "DisableForRenderFrameHostCalled"
	BackForwardCacheNotRestoredReasonDomainNotAllowed                                         BackForwardCacheNotRestoredReason = "DomainNotAllowed"
	BackForwardCacheNotRestoredReasonHTTPMethodNotGET                                         BackForwardCacheNotRestoredReason = "HTTPMethodNotGET"
	BackForwardCacheNotRestoredReasonSubframeIsNavigating                                     BackForwardCacheNotRestoredReason = "SubframeIsNavigating"
	BackForwardCacheNotRestoredReasonTimeout                                                  BackForwardCacheNotRestoredReason = "Timeout"
	BackForwardCacheNotRestoredReasonCacheLimit                                               BackForwardCacheNotRestoredReason = "CacheLimit"
	BackForwardCacheNotRestoredReasonJavaScriptExecution                                      BackForwardCacheNotRestoredReason = "JavaScriptExecution"
	BackForwardCacheNotRestoredReasonRendererProcessKilled                                    BackForwardCacheNotRestoredReason = "RendererProcessKilled"
	BackForwardCacheNotRestoredReasonRendererProcessCrashed                                   BackForwardCacheNotRestoredReason = "RendererProcessCrashed"
	BackForwardCacheNotRestoredReasonGrantedMediaStreamAccess                                 BackForwardCacheNotRestoredReason = "GrantedMediaStreamAccess"
	BackForwardCacheNotRestoredReasonSchedulerTrackedFeatureUsed                              BackForwardCacheNotRestoredReason = "SchedulerTrackedFeatureUsed"
	BackForwardCacheNotRestoredReasonConflictingBrowsingInstance                              BackForwardCacheNotRestoredReason = "ConflictingBrowsingInstance"
	BackForwardCacheNotRestoredReasonCacheFlushed                                             BackForwardCacheNotRestoredReason = "CacheFlushed"
	BackForwardCacheNotRestoredReasonServiceWorkerVersionActivation                           BackForwardCacheNotRestoredReason = "ServiceWorkerVersionActivation"
	BackForwardCacheNotRestoredReasonSessionRestored                                          BackForwardCacheNotRestoredReason = "SessionRestored"
	BackForwardCacheNotRestoredReasonServiceWorkerPostMessage                                 BackForwardCacheNotRestoredReason = "ServiceWorkerPostMessage"
	BackForwardCacheNotRestoredReasonEnteredBackForwardCacheBeforeServiceWorkerHostAdded      BackForwardCacheNotRestoredReason = "EnteredBackForwardCacheBeforeServiceWorkerHostAdded"
	BackForwardCacheNotRestoredReasonRenderFrameHostReusedSameSite                            BackForwardCacheNotRestoredReason = "RenderFrameHostReused_SameSite"
	BackForwardCacheNotRestoredReasonRenderFrameHostReusedCrossSite                           BackForwardCacheNotRestoredReason = "RenderFrameHostReused_CrossSite"
	BackForwardCacheNotRestoredReasonServiceWorkerClaim                                       BackForwardCacheNotRestoredReason = "ServiceWorkerClaim"
	BackForwardCacheNotRestoredReasonIgnoreEventAndEvict                                      BackForwardCacheNotRestoredReason = "IgnoreEventAndEvict"
	BackForwardCacheNotRestoredReasonHaveInnerContents                                        BackForwardCacheNotRestoredReason = "HaveInnerContents"
	BackForwardCacheNotRestoredReasonTimeoutPuttingInCache                                    BackForwardCacheNotRestoredReason = "TimeoutPuttingInCache"
	BackForwardCacheNotRestoredReasonBackForwardCacheDisabledByLowMemory                      BackForwardCacheNotRestoredReason = "BackForwardCacheDisabledByLowMemory"
	BackForwardCacheNotRestoredReasonBackForwardCacheDisabledByCommandLine                    BackForwardCacheNotRestoredReason = "BackForwardCacheDisabledByCommandLine"
	BackForwardCacheNotRestoredReasonNetworkRequestDatapipeDrainedAsBytesConsumer             BackForwardCacheNotRestoredReason = "NetworkRequestDatapipeDrainedAsBytesConsumer"
	BackForwardCacheNotRestoredReasonNetworkRequestRedirected                                 BackForwardCacheNotRestoredReason = "NetworkRequestRedirected"
	BackForwardCacheNotRestoredReasonNetworkRequestTimeout                                    BackForwardCacheNotRestoredReason = "NetworkRequestTimeout"
	BackForwardCacheNotRestoredReasonNetworkExceedsBufferLimit                                BackForwardCacheNotRestoredReason = "NetworkExceedsBufferLimit"
	BackForwardCacheNotRestoredReasonNavigationCancelledWhileRestoring                        BackForwardCacheNotRestoredReason = "NavigationCancelledWhileRestoring"
	BackForwardCacheNotRestoredReasonNotMostRecentNavigationEntry                             BackForwardCacheNotRestoredReason = "NotMostRecentNavigationEntry"
	BackForwardCacheNotRestoredReasonBackForwardCacheDisabledForPrerender                     BackForwardCacheNotRestoredReason = "BackForwardCacheDisabledForPrerender"
	BackForwardCacheNotRestoredReasonUserAgentOverrideDiffers                                 BackForwardCacheNotRestoredReason = "UserAgentOverrideDiffers"
	BackForwardCacheNotRestoredReasonForegroundCacheLimit                                     BackForwardCacheNotRestoredReason = "ForegroundCacheLimit"
	BackForwardCacheNotRestoredReasonBrowsingInstanceNotSwapped                               BackForwardCacheNotRestoredReason = "BrowsingInstanceNotSwapped"
	BackForwardCacheNotRestoredReasonBackForwardCacheDisabledForDelegate                      BackForwardCacheNotRestoredReason = "BackForwardCacheDisabledForDelegate"
	BackForwardCacheNotRestoredReasonOptInUnloadHeaderNotPresent                              BackForwardCacheNotRestoredReason = "OptInUnloadHeaderNotPresent"
	BackForwardCacheNotRestoredReasonUnloadHandlerExistsInSubFrame                            BackForwardCacheNotRestoredReason = "UnloadHandlerExistsInSubFrame"
	BackForwardCacheNotRestoredReasonServiceWorkerUnregistration                              BackForwardCacheNotRestoredReason = "ServiceWorkerUnregistration"
	BackForwardCacheNotRestoredReasonCacheControlNoStore                                      BackForwardCacheNotRestoredReason = "CacheControlNoStore"
	BackForwardCacheNotRestoredReasonCacheControlNoStoreCookieModified                        BackForwardCacheNotRestoredReason = "CacheControlNoStoreCookieModified"
	BackForwardCacheNotRestoredReasonCacheControlNoStoreHTTPOnlyCookieModified                BackForwardCacheNotRestoredReason = "CacheControlNoStoreHTTPOnlyCookieModified"
	BackForwardCacheNotRestoredReasonNoResponseHead                                           BackForwardCacheNotRestoredReason = "NoResponseHead"
	BackForwardCacheNotRestoredReasonUnknown                                                  BackForwardCacheNotRestoredReason = "Unknown"
	BackForwardCacheNotRestoredReasonActivationNavigationsDisallowedForBug1234857             BackForwardCacheNotRestoredReason = "ActivationNavigationsDisallowedForBug1234857"
	BackForwardCacheNotRestoredReasonWebSocket                                                BackForwardCacheNotRestoredReason = "WebSocket"
	BackForwardCacheNotRestoredReasonWebTransport                                             BackForwardCacheNotRestoredReason = "WebTransport"
	BackForwardCacheNotRestoredReasonWebRTC                                                   BackForwardCacheNotRestoredReason = "WebRTC"
	BackForwardCacheNotRestoredReasonMainResourceHasCacheControlNoStore                       BackForwardCacheNotRestoredReason = "MainResourceHasCacheControlNoStore"
	BackForwardCacheNotRestoredReasonMainResourceHasCacheControlNoCache                       BackForwardCacheNotRestoredReason = "MainResourceHasCacheControlNoCache"
	BackForwardCacheNotRestoredReasonSubresourceHasCacheControlNoStore                        BackForwardCacheNotRestoredReason = "SubresourceHasCacheControlNoStore"
	BackForwardCacheNotRestoredReasonSubresourceHasCacheControlNoCache                        BackForwardCacheNotRestoredReason = "SubresourceHasCacheControlNoCache"
	BackForwardCacheNotRestoredReasonContainsPlugins                                          BackForwardCacheNotRestoredReason = "ContainsPlugins"
	BackForwardCacheNotRestoredReasonDocumentLoaded                                           BackForwardCacheNotRestoredReason = "DocumentLoaded"
	BackForwardCacheNotRestoredReasonDedicatedWorkerOrWorklet                                 BackForwardCacheNotRestoredReason = "DedicatedWorkerOrWorklet"
	BackForwardCacheNotRestoredReasonOutstandingNetworkRequestOthers                          BackForwardCacheNotRestoredReason = "OutstandingNetworkRequestOthers"
	BackForwardCacheNotRestoredReasonOutstandingIndexedDBTransaction                          BackForwardCacheNotRestoredReason = "OutstandingIndexedDBTransaction"
	BackForwardCacheNotRestoredReasonRequestedNotificationsPermission                         BackForwardCacheNotRestoredReason = "RequestedNotificationsPermission"
	BackForwardCacheNotRestoredReasonRequestedMIDIPermission                                  BackForwardCacheNotRestoredReason = "RequestedMIDIPermission"
	BackForwardCacheNotRestoredReasonRequestedAudioCapturePermission                          BackForwardCacheNotRestoredReason = "RequestedAudioCapturePermission"
	BackForwardCacheNotRestoredReasonRequestedVideoCapturePermission                          BackForwardCacheNotRestoredReason = "RequestedVideoCapturePermission"
	BackForwardCacheNotRestoredReasonRequestedBackForwardCacheBlockedSensors                  BackForwardCacheNotRestoredReason = "RequestedBackForwardCacheBlockedSensors"
	BackForwardCacheNotRestoredReasonRequestedBackgroundWorkPermission                        BackForwardCacheNotRestoredReason = "RequestedBackgroundWorkPermission"
	BackForwardCacheNotRestoredReasonBroadcastChannel                                         BackForwardCacheNotRestoredReason = "BroadcastChannel"
	BackForwardCacheNotRestoredReasonIndexedDBConnection                                      BackForwardCacheNotRestoredReason = "IndexedDBConnection"
	BackForwardCacheNotRestoredReasonWebXR                                                    BackForwardCacheNotRestoredReason = "WebXR"
	BackForwardCacheNotRestoredReasonSharedWorker                                             BackForwardCacheNotRestoredReason = "SharedWorker"
	BackForwardCacheNotRestoredReasonWebLocks                                                 BackForwardCacheNotRestoredReason = "WebLocks"
	BackForwardCacheNotRestoredReasonWebHID                                                   BackForwardCacheNotRestoredReason = "WebHID"
	BackForwardCacheNotRestoredReasonWebShare                                                 BackForwardCacheNotRestoredReason = "WebShare"
	BackForwardCacheNotRestoredReasonRequestedStorageAccessGrant                              BackForwardCacheNotRestoredReason = "RequestedStorageAccessGrant"
	BackForwardCacheNotRestoredReasonWebNfc                                                   BackForwardCacheNotRestoredReason = "WebNfc"
	BackForwardCacheNotRestoredReasonOutstandingNetworkRequestFetch                           BackForwardCacheNotRestoredReason = "OutstandingNetworkRequestFetch"
	BackForwardCacheNotRestoredReasonOutstandingNetworkRequestXHR                             BackForwardCacheNotRestoredReason = "OutstandingNetworkRequestXHR"
	BackForwardCacheNotRestoredReasonAppBanner                                                BackForwardCacheNotRestoredReason = "AppBanner"
	BackForwardCacheNotRestoredReasonPrinting                                                 BackForwardCacheNotRestoredReason = "Printing"
	BackForwardCacheNotRestoredReasonWebDatabase                                              BackForwardCacheNotRestoredReason = "WebDatabase"
	BackForwardCacheNotRestoredReasonPictureInPicture                                         BackForwardCacheNotRestoredReason = "PictureInPicture"
	BackForwardCacheNotRestoredReasonPortal                                                   BackForwardCacheNotRestoredReason = "Portal"
	BackForwardCacheNotRestoredReasonSpeechRecognizer                                         BackForwardCacheNotRestoredReason = "SpeechRecognizer"
	BackForwardCacheNotRestoredReasonIdleManager                                              BackForwardCacheNotRestoredReason = "IdleManager"
	BackForwardCacheNotRestoredReasonPaymentManager                                           BackForwardCacheNotRestoredReason = "PaymentManager"
	BackForwardCacheNotRestoredReasonSpeechSynthesis                                          BackForwardCacheNotRestoredReason = "SpeechSynthesis"
	BackForwardCacheNotRestoredReasonKeyboardLock                                             BackForwardCacheNotRestoredReason = "KeyboardLock"
	BackForwardCacheNotRestoredReasonWebOTPService                                            BackForwardCacheNotRestoredReason = "WebOTPService"
	BackForwardCacheNotRestoredReasonOutstandingNetworkRequestDirectSocket                    BackForwardCacheNotRestoredReason = "OutstandingNetworkRequestDirectSocket"
	BackForwardCacheNotRestoredReasonIsolatedWorldScript                                      BackForwardCacheNotRestoredReason = "IsolatedWorldScript"
	BackForwardCacheNotRestoredReasonInjectedStyleSheet                                       BackForwardCacheNotRestoredReason = "InjectedStyleSheet"
	BackForwardCacheNotRestoredReasonContentSecurityHandler                                   BackForwardCacheNotRestoredReason = "ContentSecurityHandler"
	BackForwardCacheNotRestoredReasonContentWebAuthenticationAPI                              BackForwardCacheNotRestoredReason = "ContentWebAuthenticationAPI"
	BackForwardCacheNotRestoredReasonContentFileChooser                                       BackForwardCacheNotRestoredReason = "ContentFileChooser"
	BackForwardCacheNotRestoredReasonContentSerial                                            BackForwardCacheNotRestoredReason = "ContentSerial"
	BackForwardCacheNotRestoredReasonContentFileSystemAccess                                  BackForwardCacheNotRestoredReason = "ContentFileSystemAccess"
	BackForwardCacheNotRestoredReasonContentMediaDevicesDispatcherHost                        BackForwardCacheNotRestoredReason = "ContentMediaDevicesDispatcherHost"
	BackForwardCacheNotRestoredReasonContentWebBluetooth                                      BackForwardCacheNotRestoredReason = "ContentWebBluetooth"
	BackForwardCacheNotRestoredReasonContentWebUSB                                            BackForwardCacheNotRestoredReason = "ContentWebUSB"
	BackForwardCacheNotRestoredReasonContentMediaSession                                      BackForwardCacheNotRestoredReason = "ContentMediaSession"
	BackForwardCacheNotRestoredReasonEmbedderPopupBlockerTabHelper                            BackForwardCacheNotRestoredReason = "EmbedderPopupBlockerTabHelper"
	BackForwardCacheNotRestoredReasonEmbedderSafeBrowsingTriggeredPopupBlocker                BackForwardCacheNotRestoredReason = "EmbedderSafeBrowsingTriggeredPopupBlocker"
	BackForwardCacheNotRestoredReasonEmbedderSafeBrowsingThreatDetails                        BackForwardCacheNotRestoredReason = "EmbedderSafeBrowsingThreatDetails"
	BackForwardCacheNotRestoredReasonEmbedderAppBannerManager                                 BackForwardCacheNotRestoredReason = "EmbedderAppBannerManager"
	BackForwardCacheNotRestoredReasonEmbedderDomDistillerViewerSource                         BackForwardCacheNotRestoredReason = "EmbedderDomDistillerViewerSource"
	BackForwardCacheNotRestoredReasonEmbedderDomDistillerSelfDeletingRequestDelegate          BackForwardCacheNotRestoredReason = "EmbedderDomDistillerSelfDeletingRequestDelegate"
	BackForwardCacheNotRestoredReasonEmbedderOomInterventionTabHelper                         BackForwardCacheNotRestoredReason = "EmbedderOomInterventionTabHelper"
	BackForwardCacheNotRestoredReasonEmbedderOfflinePage                                      BackForwardCacheNotRestoredReason = "EmbedderOfflinePage"
	BackForwardCacheNotRestoredReasonEmbedderChromePasswordManagerClientBindCredentialManager BackForwardCacheNotRestoredReason = "EmbedderChromePasswordManagerClientBindCredentialManager"
	BackForwardCacheNotRestoredReasonEmbedderPermissionRequestManager                         BackForwardCacheNotRestoredReason = "EmbedderPermissionRequestManager"
	BackForwardCacheNotRestoredReasonEmbedderModalDialog                                      BackForwardCacheNotRestoredReason = "EmbedderModalDialog"
	BackForwardCacheNotRestoredReasonEmbedderExtensions                                       BackForwardCacheNotRestoredReason = "EmbedderExtensions"
	BackForwardCacheNotRestoredReasonEmbedderExtensionMessaging                               BackForwardCacheNotRestoredReason = "EmbedderExtensionMessaging"
	BackForwardCacheNotRestoredReasonEmbedderExtensionMessagingForOpenPort                    BackForwardCacheNotRestoredReason = "EmbedderExtensionMessagingForOpenPort"
	BackForwardCacheNotRestoredReasonEmbedderExtensionSentMessageToCachedFrame                BackForwardCacheNotRestoredReason = "EmbedderExtensionSentMessageToCachedFrame"
)

// String returns the BackForwardCacheNotRestoredReason value as a built-in string.
func (t BackForwardCacheNotRestoredReason) String() string {
	return string(t)
}

// BackForwardCacheNotRestoredReasonType data type. Types of not restored reasons for back-forward cache.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-BackForwardCacheNotRestoredReasonType
//
// This CDP type is experimental.
type BackForwardCacheNotRestoredReasonType string

// BackForwardCacheNotRestoredReasonType valid values.
const (
	BackForwardCacheNotRestoredReasonTypeSupportPending    BackForwardCacheNotRestoredReasonType = "SupportPending"
	BackForwardCacheNotRestoredReasonTypePageSupportNeeded BackForwardCacheNotRestoredReasonType = "PageSupportNeeded"
	BackForwardCacheNotRestoredReasonTypeCircumstantial    BackForwardCacheNotRestoredReasonType = "Circumstantial"
)

// String returns the BackForwardCacheNotRestoredReasonType value as a built-in string.
func (t BackForwardCacheNotRestoredReasonType) String() string {
	return string(t)
}

// BackForwardCacheNotRestoredExplanation data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-BackForwardCacheNotRestoredExplanation
//
// This CDP type is experimental.
type BackForwardCacheNotRestoredExplanation struct {
	// Type of the reason
	Type BackForwardCacheNotRestoredReasonType `json:"type"`
	// Not restored reason
	Reason BackForwardCacheNotRestoredReason `json:"reason"`
}
