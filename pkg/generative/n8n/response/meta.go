package response

type Meta struct {
	TemplateCredsSetupCompleted bool   `json:"templateCredsSetupCompleted,omitempty"`
	TemplateIdentifier          string `json:"templateId,omitempty"`
}
