package description

func New(
	title string,
	short string,
) *Description {
	return &Description{Title: title, Short: short}
}
