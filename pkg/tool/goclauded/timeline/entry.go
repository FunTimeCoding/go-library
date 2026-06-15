package timeline

type Entry struct {
	Identifier        uint
	SessionIdentifier string
	Timestamp         string
	Kind              string
	Actor             string
	Alias             string
	Full              bool
	Metadata          map[string]string
}
