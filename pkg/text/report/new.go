package report

func New(
	title string,
	maximumLength int,
) *Section {
	return &Section{title: title, maximumLength: maximumLength}
}
