package timeline

type Entry struct {
	Identifier uint
	Timestamp  string
	Kind       string
	Actor      string
	Subject    string
	Detail     string
}
