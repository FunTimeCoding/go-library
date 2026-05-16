package issue

type RequestFieldValue struct {
	FieldIdentifier string        `json:"fieldId"`
	Label           string        `json:"label"`
	Value           any           `json:"value"`
	RenderedValue   RenderedValue `json:"renderedValue,omitempty"`
}
