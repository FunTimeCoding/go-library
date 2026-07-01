package save_option

type Option struct {
	Name             string
	Content          string
	Description      string
	Type             string
	Tags             []string
	Source           string
	ParentIdentifier *int64
}
