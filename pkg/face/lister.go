package face

type Lister interface {
	List(
		collection string,
		metadata map[string]string,
		limit int,
		offset int,
		full bool,
	) (*ListOutcome, error)
}
