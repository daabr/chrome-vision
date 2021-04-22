package css

// Fires whenever a web font is updated.  A non-empty font parameter indicates a successfully loaded
// web font
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-fontsUpdated
type FontsUpdated struct {
	// The web font that has loaded.
	Font *FontFace `json:"font,omitempty"`
}

// Fires whenever a MediaQuery result changes (for example, after a browser window has been
// resized.) The current implementation considers only viewport-dependent media features.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-mediaQueryResultChanged
type MediaQueryResultChanged struct{}

// Fired whenever an active document stylesheet is added.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetAdded
type StyleSheetAdded struct {
	// Added stylesheet metainfo.
	Header CSSStyleSheetHeader
}

// Fired whenever a stylesheet is changed as a result of the client operation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetChanged
type StyleSheetChanged struct {
	StyleSheetID string
}

// Fired whenever an active document stylesheet is removed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#event-styleSheetRemoved
type StyleSheetRemoved struct {
	// Identifier of the removed stylesheet.
	StyleSheetID string
}
