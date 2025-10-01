package option

type Sed struct {
	Host       string
	Token      string
	Owner      string
	Repository string
	Branch     string
	Path       string
	Message    string
	Replaces   map[string]string
}
