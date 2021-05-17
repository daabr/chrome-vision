package audits

import (
	"github.com/daabr/chrome-vision/pkg/devtools/network"
	"github.com/daabr/chrome-vision/pkg/devtools/runtime"
)

// AffectedCookie data type. Information about a cookie that is affected by an inspector issue.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-AffectedCookie
type AffectedCookie struct {
	// The following three properties uniquely identify a cookie
	Name   string `json:"name"`
	Path   string `json:"path"`
	Domain string `json:"domain"`
}

// AffectedRequest data type. Information about a request that is affected by an inspector issue.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-AffectedRequest
type AffectedRequest struct {
	// The unique request id.
	RequestID string `json:"requestId"`
	URL       string `json:"url,omitempty"`
}

// AffectedFrame data type. Information about the frame affected by an inspector issue.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-AffectedFrame
type AffectedFrame struct {
	FrameID string `json:"frameId"`
}

// SameSiteCookieExclusionReason data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-SameSiteCookieExclusionReason
type SameSiteCookieExclusionReason string

// SameSiteCookieExclusionReason valid values.
const (
	SameSiteCookieExclusionReasonExcludeSameSiteUnspecifiedTreatedAsLax SameSiteCookieExclusionReason = "ExcludeSameSiteUnspecifiedTreatedAsLax"
	SameSiteCookieExclusionReasonExcludeSameSiteNoneInsecure            SameSiteCookieExclusionReason = "ExcludeSameSiteNoneInsecure"
	SameSiteCookieExclusionReasonExcludeSameSiteLax                     SameSiteCookieExclusionReason = "ExcludeSameSiteLax"
	SameSiteCookieExclusionReasonExcludeSameSiteStrict                  SameSiteCookieExclusionReason = "ExcludeSameSiteStrict"
)

// String returns the SameSiteCookieExclusionReason value as a built-in string.
func (t SameSiteCookieExclusionReason) String() string {
	return string(t)
}

// SameSiteCookieWarningReason data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-SameSiteCookieWarningReason
type SameSiteCookieWarningReason string

// SameSiteCookieWarningReason valid values.
const (
	SameSiteCookieWarningReasonWarnSameSiteUnspecifiedCrossSiteContext SameSiteCookieWarningReason = "WarnSameSiteUnspecifiedCrossSiteContext"
	SameSiteCookieWarningReasonWarnSameSiteNoneInsecure                SameSiteCookieWarningReason = "WarnSameSiteNoneInsecure"
	SameSiteCookieWarningReasonWarnSameSiteUnspecifiedLaxAllowUnsafe   SameSiteCookieWarningReason = "WarnSameSiteUnspecifiedLaxAllowUnsafe"
	SameSiteCookieWarningReasonWarnSameSiteStrictLaxDowngradeStrict    SameSiteCookieWarningReason = "WarnSameSiteStrictLaxDowngradeStrict"
	SameSiteCookieWarningReasonWarnSameSiteStrictCrossDowngradeStrict  SameSiteCookieWarningReason = "WarnSameSiteStrictCrossDowngradeStrict"
	SameSiteCookieWarningReasonWarnSameSiteStrictCrossDowngradeLax     SameSiteCookieWarningReason = "WarnSameSiteStrictCrossDowngradeLax"
	SameSiteCookieWarningReasonWarnSameSiteLaxCrossDowngradeStrict     SameSiteCookieWarningReason = "WarnSameSiteLaxCrossDowngradeStrict"
	SameSiteCookieWarningReasonWarnSameSiteLaxCrossDowngradeLax        SameSiteCookieWarningReason = "WarnSameSiteLaxCrossDowngradeLax"
)

// String returns the SameSiteCookieWarningReason value as a built-in string.
func (t SameSiteCookieWarningReason) String() string {
	return string(t)
}

// SameSiteCookieOperation data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-SameSiteCookieOperation
type SameSiteCookieOperation string

// SameSiteCookieOperation valid values.
const (
	SameSiteCookieOperationSetCookie  SameSiteCookieOperation = "SetCookie"
	SameSiteCookieOperationReadCookie SameSiteCookieOperation = "ReadCookie"
)

// String returns the SameSiteCookieOperation value as a built-in string.
func (t SameSiteCookieOperation) String() string {
	return string(t)
}

// SameSiteCookieIssueDetails data type. This information is currently necessary, as the front-end has a difficult
// time finding a specific cookie. With this, we can convey specific error
// information without the cookie.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-SameSiteCookieIssueDetails
type SameSiteCookieIssueDetails struct {
	Cookie                 AffectedCookie                  `json:"cookie"`
	CookieWarningReasons   []SameSiteCookieWarningReason   `json:"cookieWarningReasons"`
	CookieExclusionReasons []SameSiteCookieExclusionReason `json:"cookieExclusionReasons"`
	// Optionally identifies the site-for-cookies and the cookie url, which
	// may be used by the front-end as additional context.
	Operation      SameSiteCookieOperation `json:"operation"`
	SiteForCookies string                  `json:"siteForCookies,omitempty"`
	CookieURL      string                  `json:"cookieUrl,omitempty"`
	Request        *AffectedRequest        `json:"request,omitempty"`
}

// MixedContentResolutionStatus data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-MixedContentResolutionStatus
type MixedContentResolutionStatus string

// MixedContentResolutionStatus valid values.
const (
	MixedContentResolutionStatusMixedContentBlocked               MixedContentResolutionStatus = "MixedContentBlocked"
	MixedContentResolutionStatusMixedContentAutomaticallyUpgraded MixedContentResolutionStatus = "MixedContentAutomaticallyUpgraded"
	MixedContentResolutionStatusMixedContentWarning               MixedContentResolutionStatus = "MixedContentWarning"
)

// String returns the MixedContentResolutionStatus value as a built-in string.
func (t MixedContentResolutionStatus) String() string {
	return string(t)
}

// MixedContentResourceType data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-MixedContentResourceType
type MixedContentResourceType string

// MixedContentResourceType valid values.
const (
	MixedContentResourceTypeAudio          MixedContentResourceType = "Audio"
	MixedContentResourceTypeBeacon         MixedContentResourceType = "Beacon"
	MixedContentResourceTypeCSPReport      MixedContentResourceType = "CSPReport"
	MixedContentResourceTypeDownload       MixedContentResourceType = "Download"
	MixedContentResourceTypeEventSource    MixedContentResourceType = "EventSource"
	MixedContentResourceTypeFavicon        MixedContentResourceType = "Favicon"
	MixedContentResourceTypeFont           MixedContentResourceType = "Font"
	MixedContentResourceTypeForm           MixedContentResourceType = "Form"
	MixedContentResourceTypeFrame          MixedContentResourceType = "Frame"
	MixedContentResourceTypeImage          MixedContentResourceType = "Image"
	MixedContentResourceTypeImport         MixedContentResourceType = "Import"
	MixedContentResourceTypeManifest       MixedContentResourceType = "Manifest"
	MixedContentResourceTypePing           MixedContentResourceType = "Ping"
	MixedContentResourceTypePluginData     MixedContentResourceType = "PluginData"
	MixedContentResourceTypePluginResource MixedContentResourceType = "PluginResource"
	MixedContentResourceTypePrefetch       MixedContentResourceType = "Prefetch"
	MixedContentResourceTypeResource       MixedContentResourceType = "Resource"
	MixedContentResourceTypeScript         MixedContentResourceType = "Script"
	MixedContentResourceTypeServiceWorker  MixedContentResourceType = "ServiceWorker"
	MixedContentResourceTypeSharedWorker   MixedContentResourceType = "SharedWorker"
	MixedContentResourceTypeStylesheet     MixedContentResourceType = "Stylesheet"
	MixedContentResourceTypeTrack          MixedContentResourceType = "Track"
	MixedContentResourceTypeVideo          MixedContentResourceType = "Video"
	MixedContentResourceTypeWorker         MixedContentResourceType = "Worker"
	MixedContentResourceTypeXMLHttpRequest MixedContentResourceType = "XMLHttpRequest"
	MixedContentResourceTypeXSLT           MixedContentResourceType = "XSLT"
)

// String returns the MixedContentResourceType value as a built-in string.
func (t MixedContentResourceType) String() string {
	return string(t)
}

// MixedContentIssueDetails data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-MixedContentIssueDetails
type MixedContentIssueDetails struct {
	// The type of resource causing the mixed content issue (css, js, iframe,
	// form,...). Marked as optional because it is mapped to from
	// blink::mojom::RequestContextType, which will be replaced
	// by network::mojom::RequestDestination
	ResourceType *MixedContentResourceType `json:"resourceType,omitempty"`
	// The way the mixed content issue is being resolved.
	ResolutionStatus MixedContentResolutionStatus `json:"resolutionStatus"`
	// The unsafe http url causing the mixed content issue.
	InsecureURL string `json:"insecureURL"`
	// The url responsible for the call to an unsafe url.
	MainResourceURL string `json:"mainResourceURL"`
	// The mixed content request.
	// Does not always exist (e.g. for unsafe form submission urls).
	Request *AffectedRequest `json:"request,omitempty"`
	// Optional because not every mixed content issue is necessarily linked to a frame.
	Frame *AffectedFrame `json:"frame,omitempty"`
}

// BlockedByResponseReason data type. Enum indicating the reason a response has been blocked. These reasons are
// refinements of the net error BLOCKED_BY_RESPONSE.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-BlockedByResponseReason
type BlockedByResponseReason string

// BlockedByResponseReason valid values.
const (
	BlockedByResponseReasonCoepFrameResourceNeedsCoepHeader                  BlockedByResponseReason = "CoepFrameResourceNeedsCoepHeader"
	BlockedByResponseReasonCoopSandboxedIFrameCannotNavigateToCoopPage       BlockedByResponseReason = "CoopSandboxedIFrameCannotNavigateToCoopPage"
	BlockedByResponseReasonCorpNotSameOrigin                                 BlockedByResponseReason = "CorpNotSameOrigin"
	BlockedByResponseReasonCorpNotSameOriginAfterDefaultedToSameOriginByCoep BlockedByResponseReason = "CorpNotSameOriginAfterDefaultedToSameOriginByCoep"
	BlockedByResponseReasonCorpNotSameSite                                   BlockedByResponseReason = "CorpNotSameSite"
)

// String returns the BlockedByResponseReason value as a built-in string.
func (t BlockedByResponseReason) String() string {
	return string(t)
}

// BlockedByResponseIssueDetails data type. Details for a request that has been blocked with the BLOCKED_BY_RESPONSE
// code. Currently only used for COEP/COOP, but may be extended to include
// some CSP errors in the future.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-BlockedByResponseIssueDetails
type BlockedByResponseIssueDetails struct {
	Request      AffectedRequest         `json:"request"`
	ParentFrame  *AffectedFrame          `json:"parentFrame,omitempty"`
	BlockedFrame *AffectedFrame          `json:"blockedFrame,omitempty"`
	Reason       BlockedByResponseReason `json:"reason"`
}

// HeavyAdResolutionStatus data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-HeavyAdResolutionStatus
type HeavyAdResolutionStatus string

// HeavyAdResolutionStatus valid values.
const (
	HeavyAdResolutionStatusHeavyAdBlocked HeavyAdResolutionStatus = "HeavyAdBlocked"
	HeavyAdResolutionStatusHeavyAdWarning HeavyAdResolutionStatus = "HeavyAdWarning"
)

// String returns the HeavyAdResolutionStatus value as a built-in string.
func (t HeavyAdResolutionStatus) String() string {
	return string(t)
}

// HeavyAdReason data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-HeavyAdReason
type HeavyAdReason string

// HeavyAdReason valid values.
const (
	HeavyAdReasonNetworkTotalLimit HeavyAdReason = "NetworkTotalLimit"
	HeavyAdReasonCpuTotalLimit     HeavyAdReason = "CpuTotalLimit"
	HeavyAdReasonCpuPeakLimit      HeavyAdReason = "CpuPeakLimit"
)

// String returns the HeavyAdReason value as a built-in string.
func (t HeavyAdReason) String() string {
	return string(t)
}

// HeavyAdIssueDetails data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-HeavyAdIssueDetails
type HeavyAdIssueDetails struct {
	// The resolution status, either blocking the content or warning.
	Resolution HeavyAdResolutionStatus `json:"resolution"`
	// The reason the ad was blocked, total network or cpu or peak cpu.
	Reason HeavyAdReason `json:"reason"`
	// The frame that was blocked.
	Frame AffectedFrame `json:"frame"`
}

// ContentSecurityPolicyViolationType data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-ContentSecurityPolicyViolationType
type ContentSecurityPolicyViolationType string

// ContentSecurityPolicyViolationType valid values.
const (
	ContentSecurityPolicyViolationTypeKInlineViolation             ContentSecurityPolicyViolationType = "kInlineViolation"
	ContentSecurityPolicyViolationTypeKEvalViolation               ContentSecurityPolicyViolationType = "kEvalViolation"
	ContentSecurityPolicyViolationTypeKURLViolation                ContentSecurityPolicyViolationType = "kURLViolation"
	ContentSecurityPolicyViolationTypeKTrustedTypesSinkViolation   ContentSecurityPolicyViolationType = "kTrustedTypesSinkViolation"
	ContentSecurityPolicyViolationTypeKTrustedTypesPolicyViolation ContentSecurityPolicyViolationType = "kTrustedTypesPolicyViolation"
)

// String returns the ContentSecurityPolicyViolationType value as a built-in string.
func (t ContentSecurityPolicyViolationType) String() string {
	return string(t)
}

// SourceCodeLocation data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-SourceCodeLocation
type SourceCodeLocation struct {
	ScriptID     *runtime.ScriptID `json:"scriptId,omitempty"`
	URL          string            `json:"url"`
	LineNumber   int64             `json:"lineNumber"`
	ColumnNumber int64             `json:"columnNumber"`
}

// ContentSecurityPolicyIssueDetails data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-ContentSecurityPolicyIssueDetails
type ContentSecurityPolicyIssueDetails struct {
	// The url not included in allowed sources.
	BlockedURL string `json:"blockedURL,omitempty"`
	// Specific directive that is violated, causing the CSP issue.
	ViolatedDirective                  string                             `json:"violatedDirective"`
	IsReportOnly                       bool                               `json:"isReportOnly"`
	ContentSecurityPolicyViolationType ContentSecurityPolicyViolationType `json:"contentSecurityPolicyViolationType"`
	FrameAncestor                      *AffectedFrame                     `json:"frameAncestor,omitempty"`
	SourceCodeLocation                 *SourceCodeLocation                `json:"sourceCodeLocation,omitempty"`
	ViolatingNodeID                    int64                              `json:"violatingNodeId,omitempty"`
}

// SharedArrayBufferIssueType data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-SharedArrayBufferIssueType
type SharedArrayBufferIssueType string

// SharedArrayBufferIssueType valid values.
const (
	SharedArrayBufferIssueTypeTransferIssue SharedArrayBufferIssueType = "TransferIssue"
	SharedArrayBufferIssueTypeCreationIssue SharedArrayBufferIssueType = "CreationIssue"
)

// String returns the SharedArrayBufferIssueType value as a built-in string.
func (t SharedArrayBufferIssueType) String() string {
	return string(t)
}

// SharedArrayBufferIssueDetails data type. Details for a issue arising from an SAB being instantiated in, or
// transferred to a context that is not cross-origin isolated.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-SharedArrayBufferIssueDetails
type SharedArrayBufferIssueDetails struct {
	SourceCodeLocation SourceCodeLocation         `json:"sourceCodeLocation"`
	IsWarning          bool                       `json:"isWarning"`
	Type               SharedArrayBufferIssueType `json:"type"`
}

// TwaQualityEnforcementViolationType data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-TwaQualityEnforcementViolationType
type TwaQualityEnforcementViolationType string

// TwaQualityEnforcementViolationType valid values.
const (
	TwaQualityEnforcementViolationTypeKHttpError          TwaQualityEnforcementViolationType = "kHttpError"
	TwaQualityEnforcementViolationTypeKUnavailableOffline TwaQualityEnforcementViolationType = "kUnavailableOffline"
	TwaQualityEnforcementViolationTypeKDigitalAssetLinks  TwaQualityEnforcementViolationType = "kDigitalAssetLinks"
)

// String returns the TwaQualityEnforcementViolationType value as a built-in string.
func (t TwaQualityEnforcementViolationType) String() string {
	return string(t)
}

// TrustedWebActivityIssueDetails data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-TrustedWebActivityIssueDetails
type TrustedWebActivityIssueDetails struct {
	// The url that triggers the violation.
	URL            string                             `json:"url"`
	ViolationType  TwaQualityEnforcementViolationType `json:"violationType"`
	HttpStatusCode int64                              `json:"httpStatusCode,omitempty"`
	// The package name of the Trusted Web Activity client app. This field is
	// only used when violation type is kDigitalAssetLinks.
	PackageName string `json:"packageName,omitempty"`
	// The signature of the Trusted Web Activity client app. This field is only
	// used when violation type is kDigitalAssetLinks.
	Signature string `json:"signature,omitempty"`
}

// LowTextContrastIssueDetails data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-LowTextContrastIssueDetails
type LowTextContrastIssueDetails struct {
	ViolatingNodeID       int64   `json:"violatingNodeId"`
	ViolatingNodeSelector string  `json:"violatingNodeSelector"`
	ContrastRatio         float64 `json:"contrastRatio"`
	ThresholdAA           float64 `json:"thresholdAA"`
	ThresholdAAA          float64 `json:"thresholdAAA"`
	FontSize              string  `json:"fontSize"`
	FontWeight            string  `json:"fontWeight"`
}

// CorsIssueDetails data type. Details for a CORS related issue, e.g. a warning or error related to
// CORS RFC1918 enforcement.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-CorsIssueDetails
type CorsIssueDetails struct {
	CorsErrorStatus        network.CorsErrorStatus      `json:"corsErrorStatus"`
	IsWarning              bool                         `json:"isWarning"`
	Request                AffectedRequest              `json:"request"`
	InitiatorOrigin        string                       `json:"initiatorOrigin,omitempty"`
	ResourceIPAddressSpace *network.IPAddressSpace      `json:"resourceIPAddressSpace,omitempty"`
	ClientSecurityState    *network.ClientSecurityState `json:"clientSecurityState,omitempty"`
}

// AttributionReportingIssueType data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-AttributionReportingIssueType
type AttributionReportingIssueType string

// AttributionReportingIssueType valid values.
const (
	AttributionReportingIssueTypePermissionPolicyDisabled             AttributionReportingIssueType = "PermissionPolicyDisabled"
	AttributionReportingIssueTypeInvalidAttributionSourceEventID      AttributionReportingIssueType = "InvalidAttributionSourceEventId"
	AttributionReportingIssueTypeInvalidAttributionData               AttributionReportingIssueType = "InvalidAttributionData"
	AttributionReportingIssueTypeAttributionSourceUntrustworthyOrigin AttributionReportingIssueType = "AttributionSourceUntrustworthyOrigin"
	AttributionReportingIssueTypeAttributionUntrustworthyOrigin       AttributionReportingIssueType = "AttributionUntrustworthyOrigin"
)

// String returns the AttributionReportingIssueType value as a built-in string.
func (t AttributionReportingIssueType) String() string {
	return string(t)
}

// AttributionReportingIssueDetails data type. Details for issues around "Attribution Reporting API" usage.
// Explainer: https://github.com/WICG/conversion-measurement-api
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-AttributionReportingIssueDetails
type AttributionReportingIssueDetails struct {
	ViolationType    AttributionReportingIssueType `json:"violationType"`
	Frame            *AffectedFrame                `json:"frame,omitempty"`
	Request          *AffectedRequest              `json:"request,omitempty"`
	ViolatingNodeID  int64                         `json:"violatingNodeId,omitempty"`
	InvalidParameter string                        `json:"invalidParameter,omitempty"`
}

// InspectorIssueCode data type. A unique identifier for the type of issue. Each type may use one of the
// optional fields in InspectorIssueDetails to convey more specific
// information about the kind of issue.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-InspectorIssueCode
type InspectorIssueCode string

// InspectorIssueCode valid values.
const (
	InspectorIssueCodeSameSiteCookieIssue        InspectorIssueCode = "SameSiteCookieIssue"
	InspectorIssueCodeMixedContentIssue          InspectorIssueCode = "MixedContentIssue"
	InspectorIssueCodeBlockedByResponseIssue     InspectorIssueCode = "BlockedByResponseIssue"
	InspectorIssueCodeHeavyAdIssue               InspectorIssueCode = "HeavyAdIssue"
	InspectorIssueCodeContentSecurityPolicyIssue InspectorIssueCode = "ContentSecurityPolicyIssue"
	InspectorIssueCodeSharedArrayBufferIssue     InspectorIssueCode = "SharedArrayBufferIssue"
	InspectorIssueCodeTrustedWebActivityIssue    InspectorIssueCode = "TrustedWebActivityIssue"
	InspectorIssueCodeLowTextContrastIssue       InspectorIssueCode = "LowTextContrastIssue"
	InspectorIssueCodeCorsIssue                  InspectorIssueCode = "CorsIssue"
	InspectorIssueCodeAttributionReportingIssue  InspectorIssueCode = "AttributionReportingIssue"
)

// String returns the InspectorIssueCode value as a built-in string.
func (t InspectorIssueCode) String() string {
	return string(t)
}

// InspectorIssueDetails data type. This struct holds a list of optional fields with additional information
// specific to the kind of issue. When adding a new issue code, please also
// add a new optional field to this type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-InspectorIssueDetails
type InspectorIssueDetails struct {
	SameSiteCookieIssueDetails        *SameSiteCookieIssueDetails        `json:"sameSiteCookieIssueDetails,omitempty"`
	MixedContentIssueDetails          *MixedContentIssueDetails          `json:"mixedContentIssueDetails,omitempty"`
	BlockedByResponseIssueDetails     *BlockedByResponseIssueDetails     `json:"blockedByResponseIssueDetails,omitempty"`
	HeavyAdIssueDetails               *HeavyAdIssueDetails               `json:"heavyAdIssueDetails,omitempty"`
	ContentSecurityPolicyIssueDetails *ContentSecurityPolicyIssueDetails `json:"contentSecurityPolicyIssueDetails,omitempty"`
	SharedArrayBufferIssueDetails     *SharedArrayBufferIssueDetails     `json:"sharedArrayBufferIssueDetails,omitempty"`
	TwaQualityEnforcementDetails      *TrustedWebActivityIssueDetails    `json:"twaQualityEnforcementDetails,omitempty"`
	LowTextContrastIssueDetails       *LowTextContrastIssueDetails       `json:"lowTextContrastIssueDetails,omitempty"`
	CorsIssueDetails                  *CorsIssueDetails                  `json:"corsIssueDetails,omitempty"`
	AttributionReportingIssueDetails  *AttributionReportingIssueDetails  `json:"attributionReportingIssueDetails,omitempty"`
}

// InspectorIssue data type. An inspector issue reported from the back-end.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#type-InspectorIssue
type InspectorIssue struct {
	Code    InspectorIssueCode    `json:"code"`
	Details InspectorIssueDetails `json:"details"`
}