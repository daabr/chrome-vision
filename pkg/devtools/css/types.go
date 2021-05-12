package css

import "github.com/daabr/chrome-vision/pkg/devtools/dom"

// StyleSheetID data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-StyleSheetId
type StyleSheetID string

// StyleSheetOrigin data type. Stylesheet type: "injected" for stylesheets injected via extension, "user-agent" for user-agent
// stylesheets, "inspector" for stylesheets created by the inspector (i.e. those holding the "via
// inspector" rules), "regular" for regular stylesheets.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-StyleSheetOrigin
type StyleSheetOrigin string

// StyleSheetOrigin valid values.
const (
	StyleSheetOriginInjected  StyleSheetOrigin = "injected"
	StyleSheetOriginUserAgent StyleSheetOrigin = "user-agent"
	StyleSheetOriginInspector StyleSheetOrigin = "inspector"
	StyleSheetOriginRegular   StyleSheetOrigin = "regular"
)

// String returns the StyleSheetOrigin value as a built-in string.
func (t StyleSheetOrigin) String() string {
	return string(t)
}

// PseudoElementMatches data type. CSS rule collection for a single pseudo style.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-PseudoElementMatches
type PseudoElementMatches struct {
	// Pseudo element type.
	PseudoType dom.PseudoType `json:"pseudoType"`
	// Matches of CSS rules applicable to the pseudo style.
	Matches []RuleMatch `json:"matches"`
}

// InheritedStyleEntry data type. Inherited CSS rule collection from ancestor node.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-InheritedStyleEntry
type InheritedStyleEntry struct {
	// The ancestor node's inline style, if any, in the style inheritance chain.
	InlineStyle *CSSStyle `json:"inlineStyle,omitempty"`
	// Matches of CSS rules matching the ancestor node in the style inheritance chain.
	MatchedCSSRules []RuleMatch `json:"matchedCSSRules"`
}

// RuleMatch data type. Match data for a CSS rule.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-RuleMatch
type RuleMatch struct {
	// CSS rule in the match.
	Rule CSSRule `json:"rule"`
	// Matching selector indices in the rule's selectorList selectors (0-based).
	MatchingSelectors []int64 `json:"matchingSelectors"`
}

// Value data type. Data for a simple selector (these are delimited by commas in a selector list).
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-Value
type Value struct {
	// Value text.
	Text string `json:"text"`
	// Value range in the underlying resource (if available).
	Range *SourceRange `json:"range,omitempty"`
}

// SelectorList data type. Selector list data.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-SelectorList
type SelectorList struct {
	// Selectors in the list.
	Selectors []Value `json:"selectors"`
	// Rule selector text.
	Text string `json:"text"`
}

// CSSStyleSheetHeader data type. CSS stylesheet metainformation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSStyleSheetHeader
type CSSStyleSheetHeader struct {
	// The stylesheet identifier.
	StyleSheetID string `json:"styleSheetId"`
	// Owner frame identifier.
	FrameID string `json:"frameId"`
	// Stylesheet resource URL.
	SourceURL string `json:"sourceURL"`
	// URL of source map associated with the stylesheet (if any).
	SourceMapURL string `json:"sourceMapURL,omitempty"`
	// Stylesheet origin.
	Origin StyleSheetOrigin `json:"origin"`
	// Stylesheet title.
	Title string `json:"title"`
	// The backend id for the owner node of the stylesheet.
	OwnerNode int64 `json:"ownerNode,omitempty"`
	// Denotes whether the stylesheet is disabled.
	Disabled bool `json:"disabled"`
	// Whether the sourceURL field value comes from the sourceURL comment.
	HasSourceURL bool `json:"hasSourceURL,omitempty"`
	// Whether this stylesheet is created for STYLE tag by parser. This flag is not set for
	// document.written STYLE tags.
	IsInline bool `json:"isInline"`
	// Whether this stylesheet is mutable. Inline stylesheets become mutable
	// after they have been modified via CSSOM API.
	// <link> element's stylesheets become mutable only if DevTools modifies them.
	// Constructed stylesheets (new CSSStyleSheet()) are mutable immediately after creation.
	IsMutable bool `json:"isMutable"`
	// Whether this stylesheet is a constructed stylesheet (created using new CSSStyleSheet()).
	IsConstructed bool `json:"isConstructed"`
	// Line offset of the stylesheet within the resource (zero based).
	StartLine float64 `json:"startLine"`
	// Column offset of the stylesheet within the resource (zero based).
	StartColumn float64 `json:"startColumn"`
	// Size of the content (in characters).
	Length float64 `json:"length"`
	// Line offset of the end of the stylesheet within the resource (zero based).
	EndLine float64 `json:"endLine"`
	// Column offset of the end of the stylesheet within the resource (zero based).
	EndColumn float64 `json:"endColumn"`
}

// CSSRule data type. CSS rule representation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSRule
type CSSRule struct {
	// The css style sheet identifier (absent for user agent stylesheet and user-specified
	// stylesheet rules) this rule came from.
	StyleSheetID string `json:"styleSheetId,omitempty"`
	// Rule selector data.
	SelectorList SelectorList `json:"selectorList"`
	// Parent stylesheet's origin.
	Origin StyleSheetOrigin `json:"origin"`
	// Associated style declaration.
	Style CSSStyle `json:"style"`
	// Media list array (for rules involving media queries). The array enumerates media queries
	// starting with the innermost one, going outwards.
	Media []CSSMedia `json:"media,omitempty"`
}

// RuleUsage data type. CSS coverage information.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-RuleUsage
type RuleUsage struct {
	// The css style sheet identifier (absent for user agent stylesheet and user-specified
	// stylesheet rules) this rule came from.
	StyleSheetID string `json:"styleSheetId"`
	// Offset of the start of the rule (including selector) from the beginning of the stylesheet.
	StartOffset float64 `json:"startOffset"`
	// Offset of the end of the rule body from the beginning of the stylesheet.
	EndOffset float64 `json:"endOffset"`
	// Indicates whether the rule was actually used by some element in the page.
	Used bool `json:"used"`
}

// SourceRange data type. Text range within a resource. All numbers are zero-based.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-SourceRange
type SourceRange struct {
	// Start line of range.
	StartLine int64 `json:"startLine"`
	// Start column of range (inclusive).
	StartColumn int64 `json:"startColumn"`
	// End line of range
	EndLine int64 `json:"endLine"`
	// End column of range (exclusive).
	EndColumn int64 `json:"endColumn"`
}

// ShorthandEntry data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-ShorthandEntry
type ShorthandEntry struct {
	// Shorthand name.
	Name string `json:"name"`
	// Shorthand value.
	Value string `json:"value"`
	// Whether the property has "!important" annotation (implies `false` if absent).
	Important bool `json:"important,omitempty"`
}

// CSSComputedStyleProperty data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSComputedStyleProperty
type CSSComputedStyleProperty struct {
	// Computed style property name.
	Name string `json:"name"`
	// Computed style property value.
	Value string `json:"value"`
}

// CSSStyle data type. CSS style representation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSStyle
type CSSStyle struct {
	// The css style sheet identifier (absent for user agent stylesheet and user-specified
	// stylesheet rules) this rule came from.
	StyleSheetID string `json:"styleSheetId,omitempty"`
	// CSS properties in the style.
	CssProperties []CSSProperty `json:"cssProperties"`
	// Computed values for all shorthands found in the style.
	ShorthandEntries []ShorthandEntry `json:"shorthandEntries"`
	// Style declaration text (if available).
	CssText string `json:"cssText,omitempty"`
	// Style declaration range in the enclosing stylesheet (if available).
	Range *SourceRange `json:"range,omitempty"`
}

// CSSProperty data type. CSS property declaration data.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSProperty
type CSSProperty struct {
	// The property name.
	Name string `json:"name"`
	// The property value.
	Value string `json:"value"`
	// Whether the property has "!important" annotation (implies `false` if absent).
	Important bool `json:"important,omitempty"`
	// Whether the property is implicit (implies `false` if absent).
	Implicit bool `json:"implicit,omitempty"`
	// The full property text as specified in the style.
	Text string `json:"text,omitempty"`
	// Whether the property is understood by the browser (implies `true` if absent).
	ParsedOk bool `json:"parsedOk,omitempty"`
	// Whether the property is disabled by the user (present for source-based properties only).
	Disabled bool `json:"disabled,omitempty"`
	// The entire property range in the enclosing style declaration (if available).
	Range *SourceRange `json:"range,omitempty"`
}

// CSSMedia data type. CSS media rule descriptor.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSMedia
type CSSMedia struct {
	// Media query text.
	Text string `json:"text"`
	// Source of the media query: "mediaRule" if specified by a @media rule, "importRule" if
	// specified by an @import rule, "linkedSheet" if specified by a "media" attribute in a linked
	// stylesheet's LINK tag, "inlineSheet" if specified by a "media" attribute in an inline
	// stylesheet's STYLE tag.
	Source string `json:"source"`
	// URL of the document containing the media query description.
	SourceURL string `json:"sourceURL,omitempty"`
	// The associated rule (@media or @import) header range in the enclosing stylesheet (if
	// available).
	Range *SourceRange `json:"range,omitempty"`
	// Identifier of the stylesheet containing this object (if exists).
	StyleSheetID string `json:"styleSheetId,omitempty"`
	// Array of media queries.
	MediaList []MediaQuery `json:"mediaList,omitempty"`
}

// MediaQuery data type. Media query descriptor.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-MediaQuery
type MediaQuery struct {
	// Array of media query expressions.
	Expressions []MediaQueryExpression `json:"expressions"`
	// Whether the media query condition is satisfied.
	Active bool `json:"active"`
}

// MediaQueryExpression data type. Media query expression descriptor.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-MediaQueryExpression
type MediaQueryExpression struct {
	// Media query expression value.
	Value float64 `json:"value"`
	// Media query expression units.
	Unit string `json:"unit"`
	// Media query expression feature.
	Feature string `json:"feature"`
	// The associated range of the value text in the enclosing stylesheet (if available).
	ValueRange *SourceRange `json:"valueRange,omitempty"`
	// Computed length of media query expression (if applicable).
	ComputedLength float64 `json:"computedLength,omitempty"`
}

// PlatformFontUsage data type. Information about amount of glyphs that were rendered with given font.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-PlatformFontUsage
type PlatformFontUsage struct {
	// Font's family name reported by platform.
	FamilyName string `json:"familyName"`
	// Indicates if the font was downloaded or resolved locally.
	IsCustomFont bool `json:"isCustomFont"`
	// Amount of glyphs that were rendered with this font.
	GlyphCount float64 `json:"glyphCount"`
}

// FontVariationAxis data type. Information about font variation axes for variable fonts
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-FontVariationAxis
type FontVariationAxis struct {
	// The font-variation-setting tag (a.k.a. "axis tag").
	Tag string `json:"tag"`
	// Human-readable variation name in the default language (normally, "en").
	Name string `json:"name"`
	// The minimum value (inclusive) the font supports for this tag.
	MinValue float64 `json:"minValue"`
	// The maximum value (inclusive) the font supports for this tag.
	MaxValue float64 `json:"maxValue"`
	// The default value.
	DefaultValue float64 `json:"defaultValue"`
}

// FontFace data type. Properties of a web font: https://www.w3.org/TR/2008/REC-CSS2-20080411/fonts.html#font-descriptions
// and additional information such as platformFontFamily and fontVariationAxes.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-FontFace
type FontFace struct {
	// The font-family.
	FontFamily string `json:"fontFamily"`
	// The font-style.
	FontStyle string `json:"fontStyle"`
	// The font-variant.
	FontVariant string `json:"fontVariant"`
	// The font-weight.
	FontWeight string `json:"fontWeight"`
	// The font-stretch.
	FontStretch string `json:"fontStretch"`
	// The unicode-range.
	UnicodeRange string `json:"unicodeRange"`
	// The src.
	Src string `json:"src"`
	// The resolved platform font family
	PlatformFontFamily string `json:"platformFontFamily"`
	// Available variation settings (a.k.a. "axes").
	FontVariationAxes []FontVariationAxis `json:"fontVariationAxes,omitempty"`
}

// CSSKeyframesRule data type. CSS keyframes rule representation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSKeyframesRule
type CSSKeyframesRule struct {
	// Animation name.
	AnimationName Value `json:"animationName"`
	// List of keyframes.
	Keyframes []CSSKeyframeRule `json:"keyframes"`
}

// CSSKeyframeRule data type. CSS keyframe rule representation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-CSSKeyframeRule
type CSSKeyframeRule struct {
	// The css style sheet identifier (absent for user agent stylesheet and user-specified
	// stylesheet rules) this rule came from.
	StyleSheetID string `json:"styleSheetId,omitempty"`
	// Parent stylesheet's origin.
	Origin StyleSheetOrigin `json:"origin"`
	// Associated key text.
	KeyText Value `json:"keyText"`
	// Associated style declaration.
	Style CSSStyle `json:"style"`
}

// StyleDeclarationEdit data type. A descriptor of operation to mutate style declaration text.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#type-StyleDeclarationEdit
type StyleDeclarationEdit struct {
	// The css style sheet identifier.
	StyleSheetID string `json:"styleSheetId"`
	// The range of the style text in the enclosing stylesheet.
	Range SourceRange `json:"range"`
	// New style text.
	Text string `json:"text"`
}
