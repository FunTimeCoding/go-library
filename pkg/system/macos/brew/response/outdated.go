package response

type Outdated struct {
	Formulae []*Formula `json:"formulae"`
	Casks    []*Cask    `json:"casks"`
}
