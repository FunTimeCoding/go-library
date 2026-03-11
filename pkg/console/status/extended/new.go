package extended

func New(
	identifier int,
	name string,
	description string,
) *Extended {
	return &Extended{
		Identifier:  identifier,
		Name:        name,
		Description: description,
	}
}
