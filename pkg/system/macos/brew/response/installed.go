package response

type Installed struct {
	Formulae []InstalledFormula `json:"formulae"`
	Casks    []InstalledCask    `json:"casks"`
}
