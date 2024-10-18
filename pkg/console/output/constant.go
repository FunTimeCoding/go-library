package output

// Output formats
const (
	Text     = "text"
	Notation = "notation"
	Markdown = "markdown"
)

var (
	Default = New()
	Debug   = New().Debug()
)
