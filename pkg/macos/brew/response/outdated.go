package response

type Outdated struct {
	Formulae []string `json:"formulae"`
	Casks    []string `json:"casks"`
}
