package webauthn

// AuthenticatorID data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#type-AuthenticatorId
type AuthenticatorID string

// AuthenticatorProtocol data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#type-AuthenticatorProtocol
type AuthenticatorProtocol string

// AuthenticatorProtocol valid values.
const (
	AuthenticatorProtocolU2f   AuthenticatorProtocol = "u2f"
	AuthenticatorProtocolCtap2 AuthenticatorProtocol = "ctap2"
)

// String returns the AuthenticatorProtocol value as a built-in string.
func (t AuthenticatorProtocol) String() string {
	return string(t)
}

// Ctap2Version data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#type-Ctap2Version
type Ctap2Version string

// Ctap2Version valid values.
const (
	Ctap2VersionCtap20 Ctap2Version = "ctap2_0"
	Ctap2VersionCtap21 Ctap2Version = "ctap2_1"
)

// String returns the Ctap2Version value as a built-in string.
func (t Ctap2Version) String() string {
	return string(t)
}

// AuthenticatorTransport data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#type-AuthenticatorTransport
type AuthenticatorTransport string

// AuthenticatorTransport valid values.
const (
	AuthenticatorTransportUsb      AuthenticatorTransport = "usb"
	AuthenticatorTransportNfc      AuthenticatorTransport = "nfc"
	AuthenticatorTransportBle      AuthenticatorTransport = "ble"
	AuthenticatorTransportCable    AuthenticatorTransport = "cable"
	AuthenticatorTransportInternal AuthenticatorTransport = "internal"
)

// String returns the AuthenticatorTransport value as a built-in string.
func (t AuthenticatorTransport) String() string {
	return string(t)
}

// VirtualAuthenticatorOptions data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#type-VirtualAuthenticatorOptions
type VirtualAuthenticatorOptions struct {
	Protocol AuthenticatorProtocol `json:"protocol"`
	// Defaults to ctap2_0. Ignored if |protocol| == u2f.
	Ctap2Version *Ctap2Version          `json:"ctap2Version,omitempty"`
	Transport    AuthenticatorTransport `json:"transport"`
	// Defaults to false.
	HasResidentKey bool `json:"hasResidentKey,omitempty"`
	// Defaults to false.
	HasUserVerification bool `json:"hasUserVerification,omitempty"`
	// If set to true, the authenticator will support the largeBlob extension.
	// https://w3c.github.io/webauthn#largeBlob
	// Defaults to false.
	HasLargeBlob bool `json:"hasLargeBlob,omitempty"`
	// If set to true, the authenticator will support the credBlob extension.
	// https://fidoalliance.org/specs/fido-v2.1-rd-20201208/fido-client-to-authenticator-protocol-v2.1-rd-20201208.html#sctn-credBlob-extension
	// Defaults to false.
	HasCredBlob bool `json:"hasCredBlob,omitempty"`
	// If set to true, tests of user presence will succeed immediately.
	// Otherwise, they will not be resolved. Defaults to true.
	AutomaticPresenceSimulation bool `json:"automaticPresenceSimulation,omitempty"`
	// Sets whether User Verification succeeds or fails for an authenticator.
	// Defaults to false.
	IsUserVerified bool `json:"isUserVerified,omitempty"`
}

// Credential data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#type-Credential
type Credential struct {
	CredentialID         string `json:"credentialId"`
	IsResidentCredential bool   `json:"isResidentCredential"`
	// Relying Party ID the credential is scoped to. Must be set when adding a
	// credential.
	RpID string `json:"rpId,omitempty"`
	// The ECDSA P-256 private key in PKCS#8 format. (Encoded as a base64 string when passed over JSON)
	PrivateKey string `json:"privateKey"`
	// An opaque byte sequence with a maximum size of 64 bytes mapping the
	// credential to a specific user. (Encoded as a base64 string when passed over JSON)
	UserHandle string `json:"userHandle,omitempty"`
	// Signature counter. This is incremented by one for each successful
	// assertion.
	// See https://w3c.github.io/webauthn/#signature-counter
	SignCount int64 `json:"signCount"`
	// The large blob associated with the credential.
	// See https://w3c.github.io/webauthn/#sctn-large-blob-extension (Encoded as a base64 string when passed over JSON)
	LargeBlob string `json:"largeBlob,omitempty"`
}
