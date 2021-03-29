package webauthn

import (
	"context"
	"fmt"
)

// Enable contains the parameters, and acts as
// a Go receiver, for the CDP command `enable`.
//
// Enable the WebAuthn domain and start intercepting credential storage and
// retrieval with a virtual authenticator.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disable the WebAuthn domain.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// AddVirtualAuthenticator contains the parameters, and acts as
// a Go receiver, for the CDP command `addVirtualAuthenticator`.
//
// Creates and adds a virtual authenticator.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-addVirtualAuthenticator
type AddVirtualAuthenticator struct {
	Options VirtualAuthenticatorOptions `json:"options"`
}

// NewAddVirtualAuthenticator constructs a new AddVirtualAuthenticator struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-addVirtualAuthenticator
func NewAddVirtualAuthenticator(options VirtualAuthenticatorOptions) *AddVirtualAuthenticator {
	return &AddVirtualAuthenticator{
		Options: options,
	}
}

// AddVirtualAuthenticatorResponse contains the browser's response
// to calling the AddVirtualAuthenticator CDP command with Do().
type AddVirtualAuthenticatorResponse struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
}

// Do sends the AddVirtualAuthenticator CDP command to a browser,
// and returns the browser's response.
func (t *AddVirtualAuthenticator) Do(ctx context.Context) (*AddVirtualAuthenticatorResponse, error) {
	fmt.Println(ctx)
	return new(AddVirtualAuthenticatorResponse), nil
}

// RemoveVirtualAuthenticator contains the parameters, and acts as
// a Go receiver, for the CDP command `removeVirtualAuthenticator`.
//
// Removes the given authenticator.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-removeVirtualAuthenticator
type RemoveVirtualAuthenticator struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
}

// NewRemoveVirtualAuthenticator constructs a new RemoveVirtualAuthenticator struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-removeVirtualAuthenticator
func NewRemoveVirtualAuthenticator(authenticatorId AuthenticatorID) *RemoveVirtualAuthenticator {
	return &RemoveVirtualAuthenticator{
		AuthenticatorID: authenticatorId,
	}
}

// Do sends the RemoveVirtualAuthenticator CDP command to a browser,
// and returns the browser's response.
func (t *RemoveVirtualAuthenticator) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// AddCredential contains the parameters, and acts as
// a Go receiver, for the CDP command `addCredential`.
//
// Adds the credential to the specified authenticator.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-addCredential
type AddCredential struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
	Credential      Credential      `json:"credential"`
}

// NewAddCredential constructs a new AddCredential struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-addCredential
func NewAddCredential(authenticatorId AuthenticatorID, credential Credential) *AddCredential {
	return &AddCredential{
		AuthenticatorID: authenticatorId,
		Credential:      credential,
	}
}

// Do sends the AddCredential CDP command to a browser,
// and returns the browser's response.
func (t *AddCredential) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// GetCredential contains the parameters, and acts as
// a Go receiver, for the CDP command `getCredential`.
//
// Returns a single credential stored in the given virtual authenticator that
// matches the credential ID.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-getCredential
type GetCredential struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
	CredentialID    string          `json:"credentialId"`
}

// NewGetCredential constructs a new GetCredential struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-getCredential
func NewGetCredential(authenticatorId AuthenticatorID, credentialId string) *GetCredential {
	return &GetCredential{
		AuthenticatorID: authenticatorId,
		CredentialID:    credentialId,
	}
}

// GetCredentialResponse contains the browser's response
// to calling the GetCredential CDP command with Do().
type GetCredentialResponse struct {
	Credential Credential `json:"credential"`
}

// Do sends the GetCredential CDP command to a browser,
// and returns the browser's response.
func (t *GetCredential) Do(ctx context.Context) (*GetCredentialResponse, error) {
	fmt.Println(ctx)
	return new(GetCredentialResponse), nil
}

// GetCredentials contains the parameters, and acts as
// a Go receiver, for the CDP command `getCredentials`.
//
// Returns all the credentials stored in the given virtual authenticator.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-getCredentials
type GetCredentials struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
}

// NewGetCredentials constructs a new GetCredentials struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-getCredentials
func NewGetCredentials(authenticatorId AuthenticatorID) *GetCredentials {
	return &GetCredentials{
		AuthenticatorID: authenticatorId,
	}
}

// GetCredentialsResponse contains the browser's response
// to calling the GetCredentials CDP command with Do().
type GetCredentialsResponse struct {
	Credentials []Credential `json:"credentials"`
}

// Do sends the GetCredentials CDP command to a browser,
// and returns the browser's response.
func (t *GetCredentials) Do(ctx context.Context) (*GetCredentialsResponse, error) {
	fmt.Println(ctx)
	return new(GetCredentialsResponse), nil
}

// RemoveCredential contains the parameters, and acts as
// a Go receiver, for the CDP command `removeCredential`.
//
// Removes a credential from the authenticator.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-removeCredential
type RemoveCredential struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
	CredentialID    string          `json:"credentialId"`
}

// NewRemoveCredential constructs a new RemoveCredential struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-removeCredential
func NewRemoveCredential(authenticatorId AuthenticatorID, credentialId string) *RemoveCredential {
	return &RemoveCredential{
		AuthenticatorID: authenticatorId,
		CredentialID:    credentialId,
	}
}

// Do sends the RemoveCredential CDP command to a browser,
// and returns the browser's response.
func (t *RemoveCredential) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// ClearCredentials contains the parameters, and acts as
// a Go receiver, for the CDP command `clearCredentials`.
//
// Clears all the credentials from the specified device.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-clearCredentials
type ClearCredentials struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
}

// NewClearCredentials constructs a new ClearCredentials struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-clearCredentials
func NewClearCredentials(authenticatorId AuthenticatorID) *ClearCredentials {
	return &ClearCredentials{
		AuthenticatorID: authenticatorId,
	}
}

// Do sends the ClearCredentials CDP command to a browser,
// and returns the browser's response.
func (t *ClearCredentials) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// SetUserVerified contains the parameters, and acts as
// a Go receiver, for the CDP command `setUserVerified`.
//
// Sets whether User Verification succeeds or fails for an authenticator.
// The default is true.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-setUserVerified
type SetUserVerified struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
	IsUserVerified  bool            `json:"isUserVerified"`
}

// NewSetUserVerified constructs a new SetUserVerified struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-setUserVerified
func NewSetUserVerified(authenticatorId AuthenticatorID, isUserVerified bool) *SetUserVerified {
	return &SetUserVerified{
		AuthenticatorID: authenticatorId,
		IsUserVerified:  isUserVerified,
	}
}

// Do sends the SetUserVerified CDP command to a browser,
// and returns the browser's response.
func (t *SetUserVerified) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// SetAutomaticPresenceSimulation contains the parameters, and acts as
// a Go receiver, for the CDP command `setAutomaticPresenceSimulation`.
//
// Sets whether tests of user presence will succeed immediately (if true) or fail to resolve (if false) for an authenticator.
// The default is true.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-setAutomaticPresenceSimulation
type SetAutomaticPresenceSimulation struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
	Enabled         bool            `json:"enabled"`
}

// NewSetAutomaticPresenceSimulation constructs a new SetAutomaticPresenceSimulation struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-setAutomaticPresenceSimulation
func NewSetAutomaticPresenceSimulation(authenticatorId AuthenticatorID, enabled bool) *SetAutomaticPresenceSimulation {
	return &SetAutomaticPresenceSimulation{
		AuthenticatorID: authenticatorId,
		Enabled:         enabled,
	}
}

// Do sends the SetAutomaticPresenceSimulation CDP command to a browser,
// and returns the browser's response.
func (t *SetAutomaticPresenceSimulation) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}
