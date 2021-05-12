package headlessexperimental

// ScreenshotParams data type. Encoding options for a screenshot.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#type-ScreenshotParams
type ScreenshotParams struct {
	// Image compression format (defaults to png).
	Format string `json:"format,omitempty"`
	// Compression quality from range [0..100] (jpeg only).
	Quality int64 `json:"quality,omitempty"`
}
