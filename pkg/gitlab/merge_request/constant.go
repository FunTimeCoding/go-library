package merge_request

const (
	OpenedState = "opened"
	MergedState = "merged"
	ClosedState = "closed"
)

var States = []string{
	OpenedState,
	MergedState,
	ClosedState,
}

const OpenAlias = "open"
