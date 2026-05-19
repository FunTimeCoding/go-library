package argument

type Timeline struct {
	Since  string
	Before string
	Kind   string
	Limit  int
	Offset int
	Full   bool
}
