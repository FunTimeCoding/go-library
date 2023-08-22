package expression

func New(
	include []string,
	exclude []string,
) *Expression {
	return &Expression{include: include, exclude: exclude}
}
