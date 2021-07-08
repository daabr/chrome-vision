package css

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// AddRule contains the parameters, and acts as
// a Go receiver, for the CDP command `addRule`.
//
// Inserts a new rule with the given `ruleText` in a stylesheet with given `styleSheetId`, at the
// position specified by `location`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-addRule
type AddRule struct {
	// The css style sheet identifier where a new rule should be inserted.
	StyleSheetID string `json:"styleSheetId"`
	// The text of a new rule.
	RuleText string `json:"ruleText"`
	// Text position of a new rule in the target style sheet.
	Location SourceRange `json:"location"`
}

// NewAddRule constructs a new AddRule struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-addRule
func NewAddRule(styleSheetID string, ruleText string, location SourceRange) *AddRule {
	return &AddRule{
		StyleSheetID: styleSheetID,
		RuleText:     ruleText,
		Location:     location,
	}
}

// AddRuleResult contains the browser's response
// to calling the AddRule CDP command with Do().
type AddRuleResult struct {
	// The newly created rule.
	Rule Rule `json:"rule"`
}

// Do sends the AddRule CDP command to a browser,
// and returns the browser's response.
func (t *AddRule) Do(ctx context.Context) (*AddRuleResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.addRule", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &AddRuleResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CollectClassNames contains the parameters, and acts as
// a Go receiver, for the CDP command `collectClassNames`.
//
// Returns all class names from specified stylesheet.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-collectClassNames
type CollectClassNames struct {
	StyleSheetID string `json:"styleSheetId"`
}

// NewCollectClassNames constructs a new CollectClassNames struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-collectClassNames
func NewCollectClassNames(styleSheetID string) *CollectClassNames {
	return &CollectClassNames{
		StyleSheetID: styleSheetID,
	}
}

// CollectClassNamesResult contains the browser's response
// to calling the CollectClassNames CDP command with Do().
type CollectClassNamesResult struct {
	// Class name list.
	ClassNames []string `json:"classNames"`
}

// Do sends the CollectClassNames CDP command to a browser,
// and returns the browser's response.
func (t *CollectClassNames) Do(ctx context.Context) (*CollectClassNamesResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.collectClassNames", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &CollectClassNamesResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateStyleSheet contains the parameters, and acts as
// a Go receiver, for the CDP command `createStyleSheet`.
//
// Creates a new special "via-inspector" stylesheet in the frame with given `frameId`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-createStyleSheet
type CreateStyleSheet struct {
	// Identifier of the frame where "via-inspector" stylesheet should be created.
	FrameID string `json:"frameId"`
}

// NewCreateStyleSheet constructs a new CreateStyleSheet struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-createStyleSheet
func NewCreateStyleSheet(frameID string) *CreateStyleSheet {
	return &CreateStyleSheet{
		FrameID: frameID,
	}
}

// CreateStyleSheetResult contains the browser's response
// to calling the CreateStyleSheet CDP command with Do().
type CreateStyleSheetResult struct {
	// Identifier of the created "via-inspector" stylesheet.
	StyleSheetID string `json:"styleSheetId"`
}

// Do sends the CreateStyleSheet CDP command to a browser,
// and returns the browser's response.
func (t *CreateStyleSheet) Do(ctx context.Context) (*CreateStyleSheetResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.createStyleSheet", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &CreateStyleSheetResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables the CSS agent for the given page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	response, err := devtools.SendAndWait(ctx, "CSS.disable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// Enable contains the parameters, and acts as
// a Go receiver, for the CDP command `enable`.
//
// Enables the CSS agent for the given page. Clients should not assume that the CSS agent has been
// enabled until the result of this command is received.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	response, err := devtools.SendAndWait(ctx, "CSS.enable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// ForcePseudoState contains the parameters, and acts as
// a Go receiver, for the CDP command `forcePseudoState`.
//
// Ensures that the given node will have specified pseudo-classes whenever its style is computed by
// the browser.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-forcePseudoState
type ForcePseudoState struct {
	// The element id for which to force the pseudo state.
	NodeID int64 `json:"nodeId"`
	// Element pseudo classes to force when computing the element's style.
	ForcedPseudoClasses []string `json:"forcedPseudoClasses"`
}

// NewForcePseudoState constructs a new ForcePseudoState struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-forcePseudoState
func NewForcePseudoState(nodeID int64, forcedPseudoClasses []string) *ForcePseudoState {
	return &ForcePseudoState{
		NodeID:              nodeID,
		ForcedPseudoClasses: forcedPseudoClasses,
	}
}

// Do sends the ForcePseudoState CDP command to a browser,
// and returns the browser's response.
func (t *ForcePseudoState) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.forcePseudoState", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetBackgroundColors contains the parameters, and acts as
// a Go receiver, for the CDP command `getBackgroundColors`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getBackgroundColors
type GetBackgroundColors struct {
	// Id of the node to get background colors for.
	NodeID int64 `json:"nodeId"`
}

// NewGetBackgroundColors constructs a new GetBackgroundColors struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getBackgroundColors
func NewGetBackgroundColors(nodeID int64) *GetBackgroundColors {
	return &GetBackgroundColors{
		NodeID: nodeID,
	}
}

// GetBackgroundColorsResult contains the browser's response
// to calling the GetBackgroundColors CDP command with Do().
type GetBackgroundColorsResult struct {
	// The range of background colors behind this element, if it contains any visible text. If no
	// visible text is present, this will be undefined. In the case of a flat background color,
	// this will consist of simply that color. In the case of a gradient, this will consist of each
	// of the color stops. For anything more complicated, this will be an empty array. Images will
	// be ignored (as if the image had failed to load).
	BackgroundColors []string `json:"backgroundColors,omitempty"`
	// The computed font size for this node, as a CSS computed value string (e.g. '12px').
	ComputedFontSize string `json:"computedFontSize,omitempty"`
	// The computed font weight for this node, as a CSS computed value string (e.g. 'normal' or
	// '100').
	ComputedFontWeight string `json:"computedFontWeight,omitempty"`
}

// Do sends the GetBackgroundColors CDP command to a browser,
// and returns the browser's response.
func (t *GetBackgroundColors) Do(ctx context.Context) (*GetBackgroundColorsResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.getBackgroundColors", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetBackgroundColorsResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetComputedStyleForNode contains the parameters, and acts as
// a Go receiver, for the CDP command `getComputedStyleForNode`.
//
// Returns the computed style for a DOM node identified by `nodeId`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getComputedStyleForNode
type GetComputedStyleForNode struct {
	NodeID int64 `json:"nodeId"`
}

// NewGetComputedStyleForNode constructs a new GetComputedStyleForNode struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getComputedStyleForNode
func NewGetComputedStyleForNode(nodeID int64) *GetComputedStyleForNode {
	return &GetComputedStyleForNode{
		NodeID: nodeID,
	}
}

// GetComputedStyleForNodeResult contains the browser's response
// to calling the GetComputedStyleForNode CDP command with Do().
type GetComputedStyleForNodeResult struct {
	// Computed style for the specified DOM node.
	ComputedStyle []ComputedStyleProperty `json:"computedStyle"`
}

// Do sends the GetComputedStyleForNode CDP command to a browser,
// and returns the browser's response.
func (t *GetComputedStyleForNode) Do(ctx context.Context) (*GetComputedStyleForNodeResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.getComputedStyleForNode", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetComputedStyleForNodeResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetInlineStylesForNode contains the parameters, and acts as
// a Go receiver, for the CDP command `getInlineStylesForNode`.
//
// Returns the styles defined inline (explicitly in the "style" attribute and implicitly, using DOM
// attributes) for a DOM node identified by `nodeId`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getInlineStylesForNode
type GetInlineStylesForNode struct {
	NodeID int64 `json:"nodeId"`
}

// NewGetInlineStylesForNode constructs a new GetInlineStylesForNode struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getInlineStylesForNode
func NewGetInlineStylesForNode(nodeID int64) *GetInlineStylesForNode {
	return &GetInlineStylesForNode{
		NodeID: nodeID,
	}
}

// GetInlineStylesForNodeResult contains the browser's response
// to calling the GetInlineStylesForNode CDP command with Do().
type GetInlineStylesForNodeResult struct {
	// Inline style for the specified DOM node.
	InlineStyle *Style `json:"inlineStyle,omitempty"`
	// Attribute-defined element style (e.g. resulting from "width=20 height=100%").
	AttributesStyle *Style `json:"attributesStyle,omitempty"`
}

// Do sends the GetInlineStylesForNode CDP command to a browser,
// and returns the browser's response.
func (t *GetInlineStylesForNode) Do(ctx context.Context) (*GetInlineStylesForNodeResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.getInlineStylesForNode", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetInlineStylesForNodeResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMatchedStylesForNode contains the parameters, and acts as
// a Go receiver, for the CDP command `getMatchedStylesForNode`.
//
// Returns requested styles for a DOM node identified by `nodeId`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMatchedStylesForNode
type GetMatchedStylesForNode struct {
	NodeID int64 `json:"nodeId"`
}

// NewGetMatchedStylesForNode constructs a new GetMatchedStylesForNode struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMatchedStylesForNode
func NewGetMatchedStylesForNode(nodeID int64) *GetMatchedStylesForNode {
	return &GetMatchedStylesForNode{
		NodeID: nodeID,
	}
}

// GetMatchedStylesForNodeResult contains the browser's response
// to calling the GetMatchedStylesForNode CDP command with Do().
type GetMatchedStylesForNodeResult struct {
	// Inline style for the specified DOM node.
	InlineStyle *Style `json:"inlineStyle,omitempty"`
	// Attribute-defined element style (e.g. resulting from "width=20 height=100%").
	AttributesStyle *Style `json:"attributesStyle,omitempty"`
	// CSS rules matching this node, from all applicable stylesheets.
	MatchedCSSRules []RuleMatch `json:"matchedCSSRules,omitempty"`
	// Pseudo style matches for this node.
	PseudoElements []PseudoElementMatches `json:"pseudoElements,omitempty"`
	// A chain of inherited styles (from the immediate node parent up to the DOM tree root).
	Inherited []InheritedStyleEntry `json:"inherited,omitempty"`
	// A list of CSS keyframed animations matching this node.
	CSSKeyframesRules []KeyframesRule `json:"cssKeyframesRules,omitempty"`
}

// Do sends the GetMatchedStylesForNode CDP command to a browser,
// and returns the browser's response.
func (t *GetMatchedStylesForNode) Do(ctx context.Context) (*GetMatchedStylesForNodeResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.getMatchedStylesForNode", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetMatchedStylesForNodeResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMediaQueries contains the parameters, and acts as
// a Go receiver, for the CDP command `getMediaQueries`.
//
// Returns all media queries parsed by the rendering engine.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMediaQueries
type GetMediaQueries struct{}

// NewGetMediaQueries constructs a new GetMediaQueries struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getMediaQueries
func NewGetMediaQueries() *GetMediaQueries {
	return &GetMediaQueries{}
}

// GetMediaQueriesResult contains the browser's response
// to calling the GetMediaQueries CDP command with Do().
type GetMediaQueriesResult struct {
	Medias []Media `json:"medias"`
}

// Do sends the GetMediaQueries CDP command to a browser,
// and returns the browser's response.
func (t *GetMediaQueries) Do(ctx context.Context) (*GetMediaQueriesResult, error) {
	response, err := devtools.SendAndWait(ctx, "CSS.getMediaQueries", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetMediaQueriesResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetPlatformFontsForNode contains the parameters, and acts as
// a Go receiver, for the CDP command `getPlatformFontsForNode`.
//
// Requests information about platform fonts which we used to render child TextNodes in the given
// node.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getPlatformFontsForNode
type GetPlatformFontsForNode struct {
	NodeID int64 `json:"nodeId"`
}

// NewGetPlatformFontsForNode constructs a new GetPlatformFontsForNode struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getPlatformFontsForNode
func NewGetPlatformFontsForNode(nodeID int64) *GetPlatformFontsForNode {
	return &GetPlatformFontsForNode{
		NodeID: nodeID,
	}
}

// GetPlatformFontsForNodeResult contains the browser's response
// to calling the GetPlatformFontsForNode CDP command with Do().
type GetPlatformFontsForNodeResult struct {
	// Usage statistics for every employed platform font.
	Fonts []PlatformFontUsage `json:"fonts"`
}

// Do sends the GetPlatformFontsForNode CDP command to a browser,
// and returns the browser's response.
func (t *GetPlatformFontsForNode) Do(ctx context.Context) (*GetPlatformFontsForNodeResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.getPlatformFontsForNode", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetPlatformFontsForNodeResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetStyleSheetText contains the parameters, and acts as
// a Go receiver, for the CDP command `getStyleSheetText`.
//
// Returns the current textual content for a stylesheet.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getStyleSheetText
type GetStyleSheetText struct {
	StyleSheetID string `json:"styleSheetId"`
}

// NewGetStyleSheetText constructs a new GetStyleSheetText struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-getStyleSheetText
func NewGetStyleSheetText(styleSheetID string) *GetStyleSheetText {
	return &GetStyleSheetText{
		StyleSheetID: styleSheetID,
	}
}

// GetStyleSheetTextResult contains the browser's response
// to calling the GetStyleSheetText CDP command with Do().
type GetStyleSheetTextResult struct {
	// The stylesheet text.
	Text string `json:"text"`
}

// Do sends the GetStyleSheetText CDP command to a browser,
// and returns the browser's response.
func (t *GetStyleSheetText) Do(ctx context.Context) (*GetStyleSheetTextResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.getStyleSheetText", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetStyleSheetTextResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// TrackComputedStyleUpdates contains the parameters, and acts as
// a Go receiver, for the CDP command `trackComputedStyleUpdates`.
//
// Starts tracking the given computed styles for updates. The specified array of properties
// replaces the one previously specified. Pass empty array to disable tracking.
// Use takeComputedStyleUpdates to retrieve the list of nodes that had properties modified.
// The changes to computed style properties are only tracked for nodes pushed to the front-end
// by the DOM agent. If no changes to the tracked properties occur after the node has been pushed
// to the front-end, no updates will be issued for the node.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-trackComputedStyleUpdates
//
// This CDP method is experimental.
type TrackComputedStyleUpdates struct {
	PropertiesToTrack []ComputedStyleProperty `json:"propertiesToTrack"`
}

// NewTrackComputedStyleUpdates constructs a new TrackComputedStyleUpdates struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-trackComputedStyleUpdates
//
// This CDP method is experimental.
func NewTrackComputedStyleUpdates(propertiesToTrack []ComputedStyleProperty) *TrackComputedStyleUpdates {
	return &TrackComputedStyleUpdates{
		PropertiesToTrack: propertiesToTrack,
	}
}

// Do sends the TrackComputedStyleUpdates CDP command to a browser,
// and returns the browser's response.
func (t *TrackComputedStyleUpdates) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.trackComputedStyleUpdates", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// TakeComputedStyleUpdates contains the parameters, and acts as
// a Go receiver, for the CDP command `takeComputedStyleUpdates`.
//
// Polls the next batch of computed style updates.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-takeComputedStyleUpdates
//
// This CDP method is experimental.
type TakeComputedStyleUpdates struct{}

// NewTakeComputedStyleUpdates constructs a new TakeComputedStyleUpdates struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-takeComputedStyleUpdates
//
// This CDP method is experimental.
func NewTakeComputedStyleUpdates() *TakeComputedStyleUpdates {
	return &TakeComputedStyleUpdates{}
}

// TakeComputedStyleUpdatesResult contains the browser's response
// to calling the TakeComputedStyleUpdates CDP command with Do().
type TakeComputedStyleUpdatesResult struct {
	// The list of node Ids that have their tracked computed styles updated
	NodeIds []int64 `json:"nodeIds"`
}

// Do sends the TakeComputedStyleUpdates CDP command to a browser,
// and returns the browser's response.
func (t *TakeComputedStyleUpdates) Do(ctx context.Context) (*TakeComputedStyleUpdatesResult, error) {
	response, err := devtools.SendAndWait(ctx, "CSS.takeComputedStyleUpdates", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &TakeComputedStyleUpdatesResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetEffectivePropertyValueForNode contains the parameters, and acts as
// a Go receiver, for the CDP command `setEffectivePropertyValueForNode`.
//
// Find a rule with the given active property for the given node and set the new value for this
// property
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setEffectivePropertyValueForNode
type SetEffectivePropertyValueForNode struct {
	// The element id for which to set property.
	NodeID       int64  `json:"nodeId"`
	PropertyName string `json:"propertyName"`
	Value        string `json:"value"`
}

// NewSetEffectivePropertyValueForNode constructs a new SetEffectivePropertyValueForNode struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setEffectivePropertyValueForNode
func NewSetEffectivePropertyValueForNode(nodeID int64, propertyName string, value string) *SetEffectivePropertyValueForNode {
	return &SetEffectivePropertyValueForNode{
		NodeID:       nodeID,
		PropertyName: propertyName,
		Value:        value,
	}
}

// Do sends the SetEffectivePropertyValueForNode CDP command to a browser,
// and returns the browser's response.
func (t *SetEffectivePropertyValueForNode) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.setEffectivePropertyValueForNode", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetKeyframeKey contains the parameters, and acts as
// a Go receiver, for the CDP command `setKeyframeKey`.
//
// Modifies the keyframe rule key text.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setKeyframeKey
type SetKeyframeKey struct {
	StyleSheetID string      `json:"styleSheetId"`
	Range        SourceRange `json:"range"`
	KeyText      string      `json:"keyText"`
}

// NewSetKeyframeKey constructs a new SetKeyframeKey struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setKeyframeKey
func NewSetKeyframeKey(styleSheetID string, r SourceRange, keyText string) *SetKeyframeKey {
	return &SetKeyframeKey{
		StyleSheetID: styleSheetID,
		Range:        r,
		KeyText:      keyText,
	}
}

// SetKeyframeKeyResult contains the browser's response
// to calling the SetKeyframeKey CDP command with Do().
type SetKeyframeKeyResult struct {
	// The resulting key text after modification.
	KeyText Value `json:"keyText"`
}

// Do sends the SetKeyframeKey CDP command to a browser,
// and returns the browser's response.
func (t *SetKeyframeKey) Do(ctx context.Context) (*SetKeyframeKeyResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.setKeyframeKey", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &SetKeyframeKeyResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetMediaText contains the parameters, and acts as
// a Go receiver, for the CDP command `setMediaText`.
//
// Modifies the rule selector.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setMediaText
type SetMediaText struct {
	StyleSheetID string      `json:"styleSheetId"`
	Range        SourceRange `json:"range"`
	Text         string      `json:"text"`
}

// NewSetMediaText constructs a new SetMediaText struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setMediaText
func NewSetMediaText(styleSheetID string, r SourceRange, text string) *SetMediaText {
	return &SetMediaText{
		StyleSheetID: styleSheetID,
		Range:        r,
		Text:         text,
	}
}

// SetMediaTextResult contains the browser's response
// to calling the SetMediaText CDP command with Do().
type SetMediaTextResult struct {
	// The resulting CSS media rule after modification.
	Media Media `json:"media"`
}

// Do sends the SetMediaText CDP command to a browser,
// and returns the browser's response.
func (t *SetMediaText) Do(ctx context.Context) (*SetMediaTextResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.setMediaText", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &SetMediaTextResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetContainerQueryText contains the parameters, and acts as
// a Go receiver, for the CDP command `setContainerQueryText`.
//
// Modifies the expression of a container query.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setContainerQueryText
//
// This CDP method is experimental.
type SetContainerQueryText struct {
	StyleSheetID string      `json:"styleSheetId"`
	Range        SourceRange `json:"range"`
	Text         string      `json:"text"`
}

// NewSetContainerQueryText constructs a new SetContainerQueryText struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setContainerQueryText
//
// This CDP method is experimental.
func NewSetContainerQueryText(styleSheetID string, r SourceRange, text string) *SetContainerQueryText {
	return &SetContainerQueryText{
		StyleSheetID: styleSheetID,
		Range:        r,
		Text:         text,
	}
}

// SetContainerQueryTextResult contains the browser's response
// to calling the SetContainerQueryText CDP command with Do().
type SetContainerQueryTextResult struct {
	// The resulting CSS container query rule after modification.
	ContainerQuery ContainerQuery `json:"containerQuery"`
}

// Do sends the SetContainerQueryText CDP command to a browser,
// and returns the browser's response.
func (t *SetContainerQueryText) Do(ctx context.Context) (*SetContainerQueryTextResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.setContainerQueryText", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &SetContainerQueryTextResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetRuleSelector contains the parameters, and acts as
// a Go receiver, for the CDP command `setRuleSelector`.
//
// Modifies the rule selector.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setRuleSelector
type SetRuleSelector struct {
	StyleSheetID string      `json:"styleSheetId"`
	Range        SourceRange `json:"range"`
	Selector     string      `json:"selector"`
}

// NewSetRuleSelector constructs a new SetRuleSelector struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setRuleSelector
func NewSetRuleSelector(styleSheetID string, r SourceRange, selector string) *SetRuleSelector {
	return &SetRuleSelector{
		StyleSheetID: styleSheetID,
		Range:        r,
		Selector:     selector,
	}
}

// SetRuleSelectorResult contains the browser's response
// to calling the SetRuleSelector CDP command with Do().
type SetRuleSelectorResult struct {
	// The resulting selector list after modification.
	SelectorList SelectorList `json:"selectorList"`
}

// Do sends the SetRuleSelector CDP command to a browser,
// and returns the browser's response.
func (t *SetRuleSelector) Do(ctx context.Context) (*SetRuleSelectorResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.setRuleSelector", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &SetRuleSelectorResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetStyleSheetText contains the parameters, and acts as
// a Go receiver, for the CDP command `setStyleSheetText`.
//
// Sets the new stylesheet text.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setStyleSheetText
type SetStyleSheetText struct {
	StyleSheetID string `json:"styleSheetId"`
	Text         string `json:"text"`
}

// NewSetStyleSheetText constructs a new SetStyleSheetText struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setStyleSheetText
func NewSetStyleSheetText(styleSheetID string, text string) *SetStyleSheetText {
	return &SetStyleSheetText{
		StyleSheetID: styleSheetID,
		Text:         text,
	}
}

// SetStyleSheetTextResult contains the browser's response
// to calling the SetStyleSheetText CDP command with Do().
type SetStyleSheetTextResult struct {
	// URL of source map associated with script (if any).
	SourceMapURL string `json:"sourceMapURL,omitempty"`
}

// Do sends the SetStyleSheetText CDP command to a browser,
// and returns the browser's response.
func (t *SetStyleSheetText) Do(ctx context.Context) (*SetStyleSheetTextResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.setStyleSheetText", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &SetStyleSheetTextResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetStyleTexts contains the parameters, and acts as
// a Go receiver, for the CDP command `setStyleTexts`.
//
// Applies specified style edits one after another in the given order.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setStyleTexts
type SetStyleTexts struct {
	Edits []StyleDeclarationEdit `json:"edits"`
}

// NewSetStyleTexts constructs a new SetStyleTexts struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setStyleTexts
func NewSetStyleTexts(edits []StyleDeclarationEdit) *SetStyleTexts {
	return &SetStyleTexts{
		Edits: edits,
	}
}

// SetStyleTextsResult contains the browser's response
// to calling the SetStyleTexts CDP command with Do().
type SetStyleTextsResult struct {
	// The resulting styles after modification.
	Styles []Style `json:"styles"`
}

// Do sends the SetStyleTexts CDP command to a browser,
// and returns the browser's response.
func (t *SetStyleTexts) Do(ctx context.Context) (*SetStyleTextsResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.setStyleTexts", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &SetStyleTextsResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// StartRuleUsageTracking contains the parameters, and acts as
// a Go receiver, for the CDP command `startRuleUsageTracking`.
//
// Enables the selector recording.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-startRuleUsageTracking
type StartRuleUsageTracking struct{}

// NewStartRuleUsageTracking constructs a new StartRuleUsageTracking struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-startRuleUsageTracking
func NewStartRuleUsageTracking() *StartRuleUsageTracking {
	return &StartRuleUsageTracking{}
}

// Do sends the StartRuleUsageTracking CDP command to a browser,
// and returns the browser's response.
func (t *StartRuleUsageTracking) Do(ctx context.Context) error {
	response, err := devtools.SendAndWait(ctx, "CSS.startRuleUsageTracking", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// StopRuleUsageTracking contains the parameters, and acts as
// a Go receiver, for the CDP command `stopRuleUsageTracking`.
//
// Stop tracking rule usage and return the list of rules that were used since last call to
// `takeCoverageDelta` (or since start of coverage instrumentation)
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-stopRuleUsageTracking
type StopRuleUsageTracking struct{}

// NewStopRuleUsageTracking constructs a new StopRuleUsageTracking struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-stopRuleUsageTracking
func NewStopRuleUsageTracking() *StopRuleUsageTracking {
	return &StopRuleUsageTracking{}
}

// StopRuleUsageTrackingResult contains the browser's response
// to calling the StopRuleUsageTracking CDP command with Do().
type StopRuleUsageTrackingResult struct {
	RuleUsage []RuleUsage `json:"ruleUsage"`
}

// Do sends the StopRuleUsageTracking CDP command to a browser,
// and returns the browser's response.
func (t *StopRuleUsageTracking) Do(ctx context.Context) (*StopRuleUsageTrackingResult, error) {
	response, err := devtools.SendAndWait(ctx, "CSS.stopRuleUsageTracking", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &StopRuleUsageTrackingResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// TakeCoverageDelta contains the parameters, and acts as
// a Go receiver, for the CDP command `takeCoverageDelta`.
//
// Obtain list of rules that became used since last call to this method (or since start of coverage
// instrumentation)
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-takeCoverageDelta
type TakeCoverageDelta struct{}

// NewTakeCoverageDelta constructs a new TakeCoverageDelta struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-takeCoverageDelta
func NewTakeCoverageDelta() *TakeCoverageDelta {
	return &TakeCoverageDelta{}
}

// TakeCoverageDeltaResult contains the browser's response
// to calling the TakeCoverageDelta CDP command with Do().
type TakeCoverageDeltaResult struct {
	Coverage []RuleUsage `json:"coverage"`
	// Monotonically increasing time, in seconds.
	Timestamp float64 `json:"timestamp"`
}

// Do sends the TakeCoverageDelta CDP command to a browser,
// and returns the browser's response.
func (t *TakeCoverageDelta) Do(ctx context.Context) (*TakeCoverageDeltaResult, error) {
	response, err := devtools.SendAndWait(ctx, "CSS.takeCoverageDelta", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &TakeCoverageDeltaResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetLocalFontsEnabled contains the parameters, and acts as
// a Go receiver, for the CDP command `setLocalFontsEnabled`.
//
// Enables/disables rendering of local CSS fonts (enabled by default).
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setLocalFontsEnabled
//
// This CDP method is experimental.
type SetLocalFontsEnabled struct {
	// Whether rendering of local fonts is enabled.
	Enabled bool `json:"enabled"`
}

// NewSetLocalFontsEnabled constructs a new SetLocalFontsEnabled struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CSS/#method-setLocalFontsEnabled
//
// This CDP method is experimental.
func NewSetLocalFontsEnabled(enabled bool) *SetLocalFontsEnabled {
	return &SetLocalFontsEnabled{
		Enabled: enabled,
	}
}

// Do sends the SetLocalFontsEnabled CDP command to a browser,
// and returns the browser's response.
func (t *SetLocalFontsEnabled) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := devtools.SendAndWait(ctx, "CSS.setLocalFontsEnabled", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}
