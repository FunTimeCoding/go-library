package component

import "github.com/funtimecoding/go-library/pkg/face"

func New(w face.Worker) *Component {
	return &Component{Worker: w}
}
