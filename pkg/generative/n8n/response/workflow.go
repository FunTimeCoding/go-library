package response

import "time"

type Workflow struct {
	Create     time.Time              `json:"createdAt"`
	Update     time.Time              `json:"updatedAt"`
	Id         string                 `json:"id"`
	Name       string                 `json:"name"`
	Active     bool                   `json:"active"`
	Archived   bool                   `json:"isArchived"`
	Nodes      []*Node                `json:"nodes"`
	Connection map[string]*Connection `json:"connections"`
	Setting    struct {
		ExecutionOrder string `json:"executionOrder"`
	} `json:"settings"`
	Static any `json:"staticData"`
	Meta   struct {
		TemplateCredsSetupCompleted bool   `json:"templateCredsSetupCompleted,omitempty"`
		TemplateId                  string `json:"templateId,omitempty"`
	} `json:"meta"`
	Pin               struct{} `json:"pinData"`
	VersionIdentifier string   `json:"versionId"`
	TriggerCount      int      `json:"triggerCount"`
	Tags              []any    `json:"tags"`
}
