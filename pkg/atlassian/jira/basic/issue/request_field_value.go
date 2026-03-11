package issue

type RequestFieldValue struct {
	FieldId       string `json:"fieldId"`
	Label         string `json:"label"`
	Value         any    `json:"value"`
	RenderedValue struct {
		Html string `json:"html"`
	} `json:"renderedValue,omitempty"`
}
