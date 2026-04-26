package worker

import "github.com/funtimecoding/go-library/pkg/face"

func New(w face.Worker) *Worker {
	return &Worker{Face: w}
}
