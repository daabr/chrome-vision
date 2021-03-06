package webauthn

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
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
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "WebAuthn.enable", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Enable CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Enable) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "WebAuthn.enable", nil)
}

// ParseResponse parses the browser's response
// to the Enable CDP command.
func (t *Enable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
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
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "WebAuthn.disable", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Disable CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Disable) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "WebAuthn.disable", nil)
}

// ParseResponse parses the browser's response
// to the Disable CDP command.
func (t *Disable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
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
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-addVirtualAuthenticator
func NewAddVirtualAuthenticator(options VirtualAuthenticatorOptions) *AddVirtualAuthenticator {
	return &AddVirtualAuthenticator{
		Options: options,
	}
}

// AddVirtualAuthenticatorResult contains the browser's response
// to calling the AddVirtualAuthenticator CDP command with Do().
type AddVirtualAuthenticatorResult struct {
	AuthenticatorID string `json:"authenticatorId"`
}

// Do sends the AddVirtualAuthenticator CDP command to a browser,
// and returns the browser's response.
func (t *AddVirtualAuthenticator) Do(ctx context.Context) (*AddVirtualAuthenticatorResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "WebAuthn.addVirtualAuthenticator", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the AddVirtualAuthenticator CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *AddVirtualAuthenticator) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "WebAuthn.addVirtualAuthenticator", b)
}

// ParseResponse parses the browser's response
// to the AddVirtualAuthenticator CDP command.
func (t *AddVirtualAuthenticator) ParseResponse(m *devtools.Message) (*AddVirtualAuthenticatorResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &AddVirtualAuthenticatorResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RemoveVirtualAuthenticator contains the parameters, and acts as
// a Go receiver, for the CDP command `removeVirtualAuthenticator`.
//
// Removes the given authenticator.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-removeVirtualAuthenticator
type RemoveVirtualAuthenticator struct {
	AuthenticatorID string `json:"authenticatorId"`
}

// NewRemoveVirtualAuthenticator constructs a new RemoveVirtualAuthenticator struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-removeVirtualAuthenticator
func NewRemoveVirtualAuthenticator(authenticatorID string) *RemoveVirtualAuthenticator {
	return &RemoveVirtualAuthenticator{
		AuthenticatorID: authenticatorID,
	}
}

// Do sends the RemoveVirtualAuthenticator CDP command to a browser,
// and returns the browser's response.
func (t *RemoveVirtualAuthenticator) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "WebAuthn.removeVirtualAuthenticator", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the RemoveVirtualAuthenticator CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RemoveVirtualAuthenticator) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "WebAuthn.removeVirtualAuthenticator", b)
}

// ParseResponse parses the browser's response
// to the RemoveVirtualAuthenticator CDP command.
func (t *RemoveVirtualAuthenticator) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// AddCredential contains the parameters, and acts as
// a Go receiver, for the CDP command `addCredential`.
//
// Adds the credential to the specified authenticator.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-addCredential
type AddCredential struct {
	AuthenticatorID string     `json:"authenticatorId"`
	Credential      Credential `json:"credential"`
}

// NewAddCredential constructs a new AddCredential struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-addCredential
func NewAddCredential(authenticatorID string, credential Credential) *AddCredential {
	return &AddCredential{
		AuthenticatorID: authenticatorID,
		Credential:      credential,
	}
}

// Do sends the AddCredential CDP command to a browser,
// and returns the browser's response.
func (t *AddCredential) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "WebAuthn.addCredential", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the AddCredential CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *AddCredential) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "WebAuthn.addCredential", b)
}

// ParseResponse parses the browser's response
// to the AddCredential CDP command.
func (t *AddCredential) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
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
	AuthenticatorID string `json:"authenticatorId"`
	CredentialID    string `json:"credentialId"`
}

// NewGetCredential constructs a new GetCredential struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-getCredential
func NewGetCredential(authenticatorID string, credentialID string) *GetCredential {
	return &GetCredential{
		AuthenticatorID: authenticatorID,
		CredentialID:    credentialID,
	}
}

// GetCredentialResult contains the browser's response
// to calling the GetCredential CDP command with Do().
type GetCredentialResult struct {
	Credential Credential `json:"credential"`
}

// Do sends the GetCredential CDP command to a browser,
// and returns the browser's response.
func (t *GetCredential) Do(ctx context.Context) (*GetCredentialResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "WebAuthn.getCredential", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetCredential CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetCredential) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "WebAuthn.getCredential", b)
}

// ParseResponse parses the browser's response
// to the GetCredential CDP command.
func (t *GetCredential) ParseResponse(m *devtools.Message) (*GetCredentialResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetCredentialResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetCredentials contains the parameters, and acts as
// a Go receiver, for the CDP command `getCredentials`.
//
// Returns all the credentials stored in the given virtual authenticator.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-getCredentials
type GetCredentials struct {
	AuthenticatorID string `json:"authenticatorId"`
}

// NewGetCredentials constructs a new GetCredentials struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-getCredentials
func NewGetCredentials(authenticatorID string) *GetCredentials {
	return &GetCredentials{
		AuthenticatorID: authenticatorID,
	}
}

// GetCredentialsResult contains the browser's response
// to calling the GetCredentials CDP command with Do().
type GetCredentialsResult struct {
	Credentials []Credential `json:"credentials"`
}

// Do sends the GetCredentials CDP command to a browser,
// and returns the browser's response.
func (t *GetCredentials) Do(ctx context.Context) (*GetCredentialsResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "WebAuthn.getCredentials", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetCredentials CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetCredentials) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "WebAuthn.getCredentials", b)
}

// ParseResponse parses the browser's response
// to the GetCredentials CDP command.
func (t *GetCredentials) ParseResponse(m *devtools.Message) (*GetCredentialsResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetCredentialsResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RemoveCredential contains the parameters, and acts as
// a Go receiver, for the CDP command `removeCredential`.
//
// Removes a credential from the authenticator.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-removeCredential
type RemoveCredential struct {
	AuthenticatorID string `json:"authenticatorId"`
	CredentialID    string `json:"credentialId"`
}

// NewRemoveCredential constructs a new RemoveCredential struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-removeCredential
func NewRemoveCredential(authenticatorID string, credentialID string) *RemoveCredential {
	return &RemoveCredential{
		AuthenticatorID: authenticatorID,
		CredentialID:    credentialID,
	}
}

// Do sends the RemoveCredential CDP command to a browser,
// and returns the browser's response.
func (t *RemoveCredential) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "WebAuthn.removeCredential", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the RemoveCredential CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RemoveCredential) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "WebAuthn.removeCredential", b)
}

// ParseResponse parses the browser's response
// to the RemoveCredential CDP command.
func (t *RemoveCredential) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// ClearCredentials contains the parameters, and acts as
// a Go receiver, for the CDP command `clearCredentials`.
//
// Clears all the credentials from the specified device.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-clearCredentials
type ClearCredentials struct {
	AuthenticatorID string `json:"authenticatorId"`
}

// NewClearCredentials constructs a new ClearCredentials struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-clearCredentials
func NewClearCredentials(authenticatorID string) *ClearCredentials {
	return &ClearCredentials{
		AuthenticatorID: authenticatorID,
	}
}

// Do sends the ClearCredentials CDP command to a browser,
// and returns the browser's response.
func (t *ClearCredentials) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "WebAuthn.clearCredentials", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ClearCredentials CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ClearCredentials) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "WebAuthn.clearCredentials", b)
}

// ParseResponse parses the browser's response
// to the ClearCredentials CDP command.
func (t *ClearCredentials) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
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
	AuthenticatorID string `json:"authenticatorId"`
	IsUserVerified  bool   `json:"isUserVerified"`
}

// NewSetUserVerified constructs a new SetUserVerified struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-setUserVerified
func NewSetUserVerified(authenticatorID string, isUserVerified bool) *SetUserVerified {
	return &SetUserVerified{
		AuthenticatorID: authenticatorID,
		IsUserVerified:  isUserVerified,
	}
}

// Do sends the SetUserVerified CDP command to a browser,
// and returns the browser's response.
func (t *SetUserVerified) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "WebAuthn.setUserVerified", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetUserVerified CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetUserVerified) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "WebAuthn.setUserVerified", b)
}

// ParseResponse parses the browser's response
// to the SetUserVerified CDP command.
func (t *SetUserVerified) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
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
	AuthenticatorID string `json:"authenticatorId"`
	Enabled         bool   `json:"enabled"`
}

// NewSetAutomaticPresenceSimulation constructs a new SetAutomaticPresenceSimulation struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn/#method-setAutomaticPresenceSimulation
func NewSetAutomaticPresenceSimulation(authenticatorID string, enabled bool) *SetAutomaticPresenceSimulation {
	return &SetAutomaticPresenceSimulation{
		AuthenticatorID: authenticatorID,
		Enabled:         enabled,
	}
}

// Do sends the SetAutomaticPresenceSimulation CDP command to a browser,
// and returns the browser's response.
func (t *SetAutomaticPresenceSimulation) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "WebAuthn.setAutomaticPresenceSimulation", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetAutomaticPresenceSimulation CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetAutomaticPresenceSimulation) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "WebAuthn.setAutomaticPresenceSimulation", b)
}

// ParseResponse parses the browser's response
// to the SetAutomaticPresenceSimulation CDP command.
func (t *SetAutomaticPresenceSimulation) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}
