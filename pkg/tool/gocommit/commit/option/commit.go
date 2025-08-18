package option

type Commit struct {
	Host       string
	Token      string
	Owner      string
	Repository string
	Branch     string
	Path       string
	Message    string
	Template   string
	Replace    []string
}
