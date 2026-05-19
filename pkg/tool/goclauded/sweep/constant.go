package sweep

type action int

const (
	actionSkip action = iota
	actionCopy
	actionUpdate
)
