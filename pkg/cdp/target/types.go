package target

import (
	"github.com/daabr/chrome-vision/pkg/cdp"
	"github.com/daabr/chrome-vision/pkg/cdp/browser"
)

// https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-TargetID
type TargetID string

// Unique identifier of attached debugging session.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-SessionID
type SessionID string

// https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-TargetInfo
type TargetInfo struct {
	TargetID TargetID `json:"targetId"`
	Type     string   `json:"type"`
	Title    string   `json:"title"`
	URL      string   `json:"url"`
	// Whether the target has an attached client.
	Attached bool `json:"attached"`
	// Opener target Id
	OpenerID *TargetID `json:"openerId,omitempty"`
	// Whether the target has access to the originating window.
	//
	// This CDP property is experimental.
	CanAccessOpener bool `json:"canAccessOpener"`
	// Frame id of originating window (is only set if target has an opener).
	//
	// This CDP property is experimental.
	OpenerFrameID *cdp.FrameID `json:"openerFrameId,omitempty"`
	// This CDP property is experimental.
	BrowserContextID *browser.BrowserContextID `json:"browserContextId,omitempty"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-RemoteLocation
//
// This CDP type is experimental.
type RemoteLocation struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}
