package store

func NewSearchOption(
	query string,
	limit int,
) *SearchOption {
	return &SearchOption{
		Query: query,
		Limit: limit,
	}
}
