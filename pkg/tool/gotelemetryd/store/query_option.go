package store

type QueryOption struct {
	Tool    string
	Surface string
	Actor   string
	Since   string
	Until   string
	Limit   int
	Offset  int
}
