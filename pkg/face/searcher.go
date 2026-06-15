package face

type Searcher interface {
	Search(
		query string,
		collection string,
		limit int,
	) ([]SearchResult, error)
}
