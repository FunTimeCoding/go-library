package issue

type RequestFieldValue struct {
	FieldId       string `json:"fieldId"`
	Label         string `json:"label"`
	Value         any    `json:"value"`
	RenderedValue RenderedValue `json:"renderedValue,omitempty"`
}
