package response

type Meta struct {
	TemplateCredsSetupCompleted bool   `json:"templateCredsSetupCompleted,omitempty"`
	TemplateId                  string `json:"templateId,omitempty"`
}
