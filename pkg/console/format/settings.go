package format

type Settings struct {
	UseColor bool

	// Extended Additional multiple lines of information
	ShowExtended bool

	// Tags Toggles for fine-grained control of what to display
	Tags []string

	// Raw Print the wrapped object with %+v in a new line
	ShowRaw bool
}
