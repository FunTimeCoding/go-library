package output

import "github.com/funtimecoding/go-library/pkg/lint/concern"

type Results struct {
	workDirectory string
	Entries       []*concern.Concern
}
