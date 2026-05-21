package resolve_result

func NewMatch(
	identifier string,
	name string,
	alias string,
	field string,
) *Match {
	return &Match{
		Identifier: identifier,
		Name:       name,
		Alias:      alias,
		Field:      field,
	}
}
