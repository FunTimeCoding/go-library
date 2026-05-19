package notation

import "encoding/json"

type Line struct {
	Type      string          `json:"type"`
	Slug      string          `json:"slug"`
	Timestamp string          `json:"timestamp"`
	SessionID string          `json:"sessionId"`
	CWD       string          `json:"cwd"`
	GitBranch string          `json:"gitBranch"`
	IsMeta    bool            `json:"isMeta"`
	Message   json.RawMessage `json:"message"`
}
