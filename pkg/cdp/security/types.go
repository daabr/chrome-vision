package security

// CertificateID data type. An internal certificate ID value.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-CertificateId
type CertificateID int64

// MixedContentType data type. A description of mixed content (HTTP resources on HTTPS pages), as defined by
// https://www.w3.org/TR/mixed-content/#categories
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-MixedContentType
type MixedContentType string

// MixedContentType valid values.
const (
	MixedContentTypeBlockable           MixedContentType = "blockable"
	MixedContentTypeOptionallyBlockable MixedContentType = "optionally-blockable"
	MixedContentTypeNone                MixedContentType = "none"
)

// SecurityState data type. The security level of a page or resource.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-SecurityState
type SecurityState string

// SecurityState valid values.
const (
	SecurityStateUnknown        SecurityState = "unknown"
	SecurityStateNeutral        SecurityState = "neutral"
	SecurityStateInsecure       SecurityState = "insecure"
	SecurityStateSecure         SecurityState = "secure"
	SecurityStateInfo           SecurityState = "info"
	SecurityStateInsecureBroken SecurityState = "insecure-broken"
)

// CertificateSecurityState data type. Details about the security state of the page certificate.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-CertificateSecurityState
//
// This CDP type is experimental.
type CertificateSecurityState struct {
	// Protocol name (e.g. "TLS 1.2" or "QUIC").
	Protocol string `json:"protocol"`
	// Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchange string `json:"keyExchange"`
	// (EC)DH group used by the connection, if applicable.
	KeyExchangeGroup string `json:"keyExchangeGroup,omitempty"`
	// Cipher name.
	Cipher string `json:"cipher"`
	// TLS MAC. Note that AEAD ciphers do not have separate MACs.
	Mac string `json:"mac,omitempty"`
	// Page certificate.
	Certificate []string `json:"certificate"`
	// Certificate subject name.
	SubjectName string `json:"subjectName"`
	// Name of the issuing CA.
	Issuer string `json:"issuer"`
	// Certificate valid from date.
	ValidFrom float64 `json:"validFrom"`
	// Certificate valid to (expiration) date
	ValidTo float64 `json:"validTo"`
	// The highest priority network error code, if the certificate has an error.
	CertificateNetworkError string `json:"certificateNetworkError,omitempty"`
	// True if the certificate uses a weak signature aglorithm.
	CertificateHasWeakSignature bool `json:"certificateHasWeakSignature"`
	// True if the certificate has a SHA1 signature in the chain.
	CertificateHasSha1Signature bool `json:"certificateHasSha1Signature"`
	// True if modern SSL
	ModernSSL bool `json:"modernSSL"`
	// True if the connection is using an obsolete SSL protocol.
	ObsoleteSslProtocol bool `json:"obsoleteSslProtocol"`
	// True if the connection is using an obsolete SSL key exchange.
	ObsoleteSslKeyExchange bool `json:"obsoleteSslKeyExchange"`
	// True if the connection is using an obsolete SSL cipher.
	ObsoleteSslCipher bool `json:"obsoleteSslCipher"`
	// True if the connection is using an obsolete SSL signature.
	ObsoleteSslSignature bool `json:"obsoleteSslSignature"`
}

// SafetyTipStatus data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-SafetyTipStatus
//
// This CDP type is experimental.
type SafetyTipStatus string

// SafetyTipStatus valid values.
const (
	SafetyTipStatusBadReputation SafetyTipStatus = "badReputation"
	SafetyTipStatusLookalike     SafetyTipStatus = "lookalike"
)

// SafetyTipInfo data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-SafetyTipInfo
//
// This CDP type is experimental.
type SafetyTipInfo struct {
	// Describes whether the page triggers any safety tips or reputation warnings. Default is unknown.
	SafetyTipStatus SafetyTipStatus `json:"safetyTipStatus"`
	// The URL the safety tip suggested ("Did you mean?"). Only filled in for lookalike matches.
	SafeURL string `json:"safeUrl,omitempty"`
}

// VisibleSecurityState data type. Security state information about the page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-VisibleSecurityState
//
// This CDP type is experimental.
type VisibleSecurityState struct {
	// The security level of the page.
	SecurityState SecurityState `json:"securityState"`
	// Security state details about the page certificate.
	CertificateSecurityState *CertificateSecurityState `json:"certificateSecurityState,omitempty"`
	// The type of Safety Tip triggered on the page. Note that this field will be set even if the Safety Tip UI was not actually shown.
	SafetyTipInfo *SafetyTipInfo `json:"safetyTipInfo,omitempty"`
	// Array of security state issues ids.
	SecurityStateIssueIds []string `json:"securityStateIssueIds"`
}

// SecurityStateExplanation data type. An explanation of an factor contributing to the security state.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-SecurityStateExplanation
type SecurityStateExplanation struct {
	// Security state representing the severity of the factor being explained.
	SecurityState SecurityState `json:"securityState"`
	// Title describing the type of factor.
	Title string `json:"title"`
	// Short phrase describing the type of factor.
	Summary string `json:"summary"`
	// Full text explanation of the factor.
	Description string `json:"description"`
	// The type of mixed content described by the explanation.
	MixedContentType MixedContentType `json:"mixedContentType"`
	// Page certificate.
	Certificate []string `json:"certificate"`
	// Recommendations to fix any issues.
	Recommendations []string `json:"recommendations,omitempty"`
}

// InsecureContentStatus data type. Information about insecure content on the page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-InsecureContentStatus
//
// This CDP type is deprecated.
type InsecureContentStatus struct {
	// Always false.
	RanMixedContent bool `json:"ranMixedContent"`
	// Always false.
	DisplayedMixedContent bool `json:"displayedMixedContent"`
	// Always false.
	ContainedMixedForm bool `json:"containedMixedForm"`
	// Always false.
	RanContentWithCertErrors bool `json:"ranContentWithCertErrors"`
	// Always false.
	DisplayedContentWithCertErrors bool `json:"displayedContentWithCertErrors"`
	// Always set to unknown.
	RanInsecureContentStyle SecurityState `json:"ranInsecureContentStyle"`
	// Always set to unknown.
	DisplayedInsecureContentStyle SecurityState `json:"displayedInsecureContentStyle"`
}

// CertificateErrorAction data type. The action to take when a certificate error occurs. continue will continue processing the
// request and cancel will cancel the request.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#type-CertificateErrorAction
type CertificateErrorAction string

// CertificateErrorAction valid values.
const (
	CertificateErrorActionContinue CertificateErrorAction = "continue"
	CertificateErrorActionCancel   CertificateErrorAction = "cancel"
)
