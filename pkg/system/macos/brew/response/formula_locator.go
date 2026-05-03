package response

type FormulaLocator struct {
	Stable StableSource `json:"stable"`
	Head   HeadSource   `json:"head,omitempty"`
}
