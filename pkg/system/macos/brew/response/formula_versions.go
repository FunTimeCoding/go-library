package response

type FormulaVersions struct {
	Stable string  `json:"stable"`
	Head   *string `json:"head"`
	Bottle bool    `json:"bottle"`
}
