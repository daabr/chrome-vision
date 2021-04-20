package cdp

import (
	"context"
	"encoding/json"
)

// Partial copy of target.TargetInfo.
type targetInfo struct {
	TargetID string `json:"targetId"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	Attached bool   `json:"attached"`
}

type targetEvent struct {
	TargetInfo targetInfo `json:"targetInfo"`
}

// Parse incoming CDP target state events to update the session's state.
// Called as a goroutine in session.go, before calling "Target.setDiscoverTargets".
func receiveTargetUpdates(ctx context.Context, channels []chan *Message) {
	session, _ := FromContext(ctx)
	for session.browserDone != nil {
		select {
		case m := <-channels[0]: // Target.targetCreated (targetEvent).
			t := &targetEvent{}
			json.Unmarshal(m.Params, t)
			session.targets[t.TargetInfo.TargetID] = t.TargetInfo
		case m := <-channels[1]: // Target.targetInfoChanged (targetEvent).
			t := &targetEvent{}
			json.Unmarshal(m.Params, t)
			session.targets[t.TargetInfo.TargetID] = t.TargetInfo
		case m := <-channels[2]: // Target.targetDestroyed (targetInfo).
			t := &targetInfo{}
			json.Unmarshal(m.Params, t)
			delete(session.targets, t.TargetID)
		}
	}
}