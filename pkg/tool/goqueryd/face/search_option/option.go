package search_option

type Option struct {
	Query      string
	Collection string
	Limit      int
	Exclude    []string
}
