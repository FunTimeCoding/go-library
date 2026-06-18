package argument

type AddImport struct {
	File       string `json:"file"`
	ImportPath string `json:"import_path"`
	Alias      string `json:"alias"`
}
