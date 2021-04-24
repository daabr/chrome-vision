package audits

// IssueAdded asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#event-issueAdded
type IssueAdded struct {
	Issue InspectorIssue `json:"issue"`
}
