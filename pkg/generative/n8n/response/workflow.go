package response

import "time"

type Workflow struct {
	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
	Id          string                 `json:"id"`
	Name        string                 `json:"name"`
	Active      bool                   `json:"active"`
	IsArchived  bool                   `json:"isArchived"`
	Nodes       []*Node                `json:"nodes"`
	Connections map[string]*Connection `json:"connections"`
	Settings    struct {
		ExecutionOrder string `json:"executionOrder"`
	} `json:"settings"`
	StaticData any `json:"staticData"`
	Meta       struct {
		TemplateCredsSetupCompleted bool   `json:"templateCredsSetupCompleted,omitempty"`
		TemplateId                  string `json:"templateId,omitempty"`
	} `json:"meta"`
	PinData struct {
	} `json:"pinData"`
	VersionId    string `json:"versionId"`
	TriggerCount int    `json:"triggerCount"`
	Tags         []any  `json:"tags"`
}
