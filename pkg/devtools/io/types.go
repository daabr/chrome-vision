package io

// StreamHandle data type. This is either obtained from another method or specified as `blob:&lt;uuid&gt;` where
// `&lt;uuid&gt` is an UUID of a Blob.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IO/#type-StreamHandle
type StreamHandle string
