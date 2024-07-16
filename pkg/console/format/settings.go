package format

type Settings struct {
	Color bool

	// Extended Additional multiple lines of information
	Extended bool

	// Tags Toggles for fine-grained control of what to display
	Tags []string

	// Raw Print the wrapped object with %+v in a new line
	Raw bool
}
