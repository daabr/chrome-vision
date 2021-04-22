package input

// Emitted only when `Input.setInterceptDrags` is enabled. Use this data with `Input.dispatchDragEvent` to
// restore normal drag and drop behavior.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Input/#event-dragIntercepted
//
// This CDP event is experimental.
type DragIntercepted struct {
	Data DragData
}
