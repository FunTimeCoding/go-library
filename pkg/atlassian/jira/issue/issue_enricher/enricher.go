package issue_enricher

type Enricher struct {
	concerns          Slice
	score             Float
	commentNameFilter []string
}
