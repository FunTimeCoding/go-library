package store

type QueryOption struct {
	Tool    string
	Surface string
	Actor   string
	Kind    string
	Since   string
	Until   string
	Limit   int
	Offset  int
}
