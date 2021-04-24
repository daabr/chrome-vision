package cast

// SinksUpdated asynchronous event. This is fired whenever the list of available sinks changes. A sink is a
// device or a software surface that you can cast to.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Cast/#event-sinksUpdated
type SinksUpdated struct {
	Sinks []Sink `json:"sinks"`
}

// IssueUpdated asynchronous event. This is fired whenever the outstanding issue/error message changes.
// |issueMessage| is empty if there is no issue.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Cast/#event-issueUpdated
type IssueUpdated struct {
	IssueMessage string `json:"issueMessage"`
}
