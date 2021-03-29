package accessibility

import (
	"encoding/json"

	"github.com/daabr/chrome-vision/pkg/cdp/dom"
)

// Unique accessibility node identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXNodeId
type AXNodeID string

// Enum of possible property types.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXValueType
type AXValueType string

// AXValueType valid values.
const (
	AXValueTypeBoolean            AXValueType = "boolean"
	AXValueTypeTristate           AXValueType = "tristate"
	AXValueTypeBooleanOrUndefined AXValueType = "booleanOrUndefined"
	AXValueTypeIdref              AXValueType = "idref"
	AXValueTypeIdrefList          AXValueType = "idrefList"
	AXValueTypeInteger            AXValueType = "integer"
	AXValueTypeNode               AXValueType = "node"
	AXValueTypeNodeList           AXValueType = "nodeList"
	AXValueTypeNumber             AXValueType = "number"
	AXValueTypeString             AXValueType = "string"
	AXValueTypeComputedString     AXValueType = "computedString"
	AXValueTypeToken              AXValueType = "token"
	AXValueTypeTokenList          AXValueType = "tokenList"
	AXValueTypeDomRelation        AXValueType = "domRelation"
	AXValueTypeRole               AXValueType = "role"
	AXValueTypeInternalRole       AXValueType = "internalRole"
	AXValueTypeValueUndefined     AXValueType = "valueUndefined"
)

// Enum of possible property sources.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXValueSourceType
type AXValueSourceType string

// AXValueSourceType valid values.
const (
	AXValueSourceTypeAttribute      AXValueSourceType = "attribute"
	AXValueSourceTypeImplicit       AXValueSourceType = "implicit"
	AXValueSourceTypeStyle          AXValueSourceType = "style"
	AXValueSourceTypeContents       AXValueSourceType = "contents"
	AXValueSourceTypePlaceholder    AXValueSourceType = "placeholder"
	AXValueSourceTypeRelatedElement AXValueSourceType = "relatedElement"
)

// Enum of possible native property sources (as a subtype of a particular AXValueSourceType).
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXValueNativeSourceType
type AXValueNativeSourceType string

// AXValueNativeSourceType valid values.
const (
	AXValueNativeSourceTypeFigcaption     AXValueNativeSourceType = "figcaption"
	AXValueNativeSourceTypeLabel          AXValueNativeSourceType = "label"
	AXValueNativeSourceTypeLabelfor       AXValueNativeSourceType = "labelfor"
	AXValueNativeSourceTypeLabelwrapped   AXValueNativeSourceType = "labelwrapped"
	AXValueNativeSourceTypeLegend         AXValueNativeSourceType = "legend"
	AXValueNativeSourceTypeRubyannotation AXValueNativeSourceType = "rubyannotation"
	AXValueNativeSourceTypeTablecaption   AXValueNativeSourceType = "tablecaption"
	AXValueNativeSourceTypeTitle          AXValueNativeSourceType = "title"
	AXValueNativeSourceTypeOther          AXValueNativeSourceType = "other"
)

// A single source for a computed AX property.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXValueSource
type AXValueSource struct {
	// What type of source this is.
	Type AXValueSourceType `json:"type"`
	// The value of this property source.
	Value *AXValue `json:"value,omitempty"`
	// The name of the relevant attribute, if any.
	Attribute string `json:"attribute,omitempty"`
	// The value of the relevant attribute, if any.
	AttributeValue *AXValue `json:"attributeValue,omitempty"`
	// Whether this source is superseded by a higher priority source.
	Superseded bool `json:"superseded,omitempty"`
	// The native markup source for this value, e.g. a <label> element.
	NativeSource *AXValueNativeSourceType `json:"nativeSource,omitempty"`
	// The value, such as a node or node list, of the native source.
	NativeSourceValue *AXValue `json:"nativeSourceValue,omitempty"`
	// Whether the value for this property is invalid.
	Invalid bool `json:"invalid,omitempty"`
	// Reason for the value being invalid, if it is.
	InvalidReason string `json:"invalidReason,omitempty"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXRelatedNode
type AXRelatedNode struct {
	// The BackendNodeId of the related DOM node.
	BackendDOMNodeID dom.BackendNodeID `json:"backendDOMNodeId"`
	// The IDRef value provided, if any.
	Idref string `json:"idref,omitempty"`
	// The text alternative of this node in the current context.
	Text string `json:"text,omitempty"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXProperty
type AXProperty struct {
	// The name of this property.
	Name AXPropertyName `json:"name"`
	// The value of this property.
	Value AXValue `json:"value"`
}

// A single computed AX property.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXValue
type AXValue struct {
	// The type of this value.
	Type AXValueType `json:"type"`
	// The computed value of this property.
	Value json.RawMessage `json:"value,omitempty"`
	// One or more related nodes, if applicable.
	RelatedNodes []AXRelatedNode `json:"relatedNodes,omitempty"`
	// The sources which contributed to the computation of this property.
	Sources []AXValueSource `json:"sources,omitempty"`
}

// Values of AXProperty name:
// - from 'busy' to 'roledescription': states which apply to every AX node
// - from 'live' to 'root': attributes which apply to nodes in live regions
// - from 'autocomplete' to 'valuetext': attributes which apply to widgets
// - from 'checked' to 'selected': states which apply to widgets
// - from 'activedescendant' to 'owns' - relationships between elements other than parent/child/sibling.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXPropertyName
type AXPropertyName string

// AXPropertyName valid values.
const (
	AXPropertyNameBusy             AXPropertyName = "busy"
	AXPropertyNameDisabled         AXPropertyName = "disabled"
	AXPropertyNameEditable         AXPropertyName = "editable"
	AXPropertyNameFocusable        AXPropertyName = "focusable"
	AXPropertyNameFocused          AXPropertyName = "focused"
	AXPropertyNameHidden           AXPropertyName = "hidden"
	AXPropertyNameHiddenRoot       AXPropertyName = "hiddenRoot"
	AXPropertyNameInvalid          AXPropertyName = "invalid"
	AXPropertyNameKeyshortcuts     AXPropertyName = "keyshortcuts"
	AXPropertyNameSettable         AXPropertyName = "settable"
	AXPropertyNameRoledescription  AXPropertyName = "roledescription"
	AXPropertyNameLive             AXPropertyName = "live"
	AXPropertyNameAtomic           AXPropertyName = "atomic"
	AXPropertyNameRelevant         AXPropertyName = "relevant"
	AXPropertyNameRoot             AXPropertyName = "root"
	AXPropertyNameAutocomplete     AXPropertyName = "autocomplete"
	AXPropertyNameHasPopup         AXPropertyName = "hasPopup"
	AXPropertyNameLevel            AXPropertyName = "level"
	AXPropertyNameMultiselectable  AXPropertyName = "multiselectable"
	AXPropertyNameOrientation      AXPropertyName = "orientation"
	AXPropertyNameMultiline        AXPropertyName = "multiline"
	AXPropertyNameReadonly         AXPropertyName = "readonly"
	AXPropertyNameRequired         AXPropertyName = "required"
	AXPropertyNameValuemin         AXPropertyName = "valuemin"
	AXPropertyNameValuemax         AXPropertyName = "valuemax"
	AXPropertyNameValuetext        AXPropertyName = "valuetext"
	AXPropertyNameChecked          AXPropertyName = "checked"
	AXPropertyNameExpanded         AXPropertyName = "expanded"
	AXPropertyNameModal            AXPropertyName = "modal"
	AXPropertyNamePressed          AXPropertyName = "pressed"
	AXPropertyNameSelected         AXPropertyName = "selected"
	AXPropertyNameActivedescendant AXPropertyName = "activedescendant"
	AXPropertyNameControls         AXPropertyName = "controls"
	AXPropertyNameDescribedby      AXPropertyName = "describedby"
	AXPropertyNameDetails          AXPropertyName = "details"
	AXPropertyNameErrormessage     AXPropertyName = "errormessage"
	AXPropertyNameFlowto           AXPropertyName = "flowto"
	AXPropertyNameLabelledby       AXPropertyName = "labelledby"
	AXPropertyNameOwns             AXPropertyName = "owns"
)

// A node in the accessibility tree.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#type-AXNode
type AXNode struct {
	// Unique identifier for this node.
	NodeID AXNodeID `json:"nodeId"`
	// Whether this node is ignored for accessibility
	Ignored bool `json:"ignored"`
	// Collection of reasons why this node is hidden.
	IgnoredReasons []AXProperty `json:"ignoredReasons,omitempty"`
	// This `Node`'s role, whether explicit or implicit.
	Role *AXValue `json:"role,omitempty"`
	// The accessible name for this `Node`.
	Name *AXValue `json:"name,omitempty"`
	// The accessible description for this `Node`.
	Description *AXValue `json:"description,omitempty"`
	// The value for this `Node`.
	Value *AXValue `json:"value,omitempty"`
	// All other properties
	Properties []AXProperty `json:"properties,omitempty"`
	// IDs for each of this node's child nodes.
	ChildIds []AXNodeID `json:"childIds,omitempty"`
	// The backend ID for the associated DOM node, if any.
	BackendDOMNodeID *dom.BackendNodeID `json:"backendDOMNodeId,omitempty"`
}
