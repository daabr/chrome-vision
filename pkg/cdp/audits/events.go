package audits

// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#event-issueAdded
type IssueAdded struct {
	Issue InspectorIssue
}
