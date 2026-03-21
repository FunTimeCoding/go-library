package issue

type RequestFieldValue struct {
	FieldId       string `json:"fieldId"`
	Label         string `json:"label"`
	Value         any    `json:"value"`
	RenderedValue struct {
		Markup string `json:"html"`
	} `json:"renderedValue,omitempty"`
}
