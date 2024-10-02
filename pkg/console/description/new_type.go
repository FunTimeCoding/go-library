package description

func NewType(
	resourceType string,
	title string,
	short string,
) *Description {
	result := New(title, short)
	result.Type = resourceType

	return result
}
