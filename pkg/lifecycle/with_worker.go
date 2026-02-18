package lifecycle

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle/component"
)

func WithWorker(w face.Worker) Option {
	return func(l *Lifecycle) {
		l.component = append(
			l.component,
			&component.Component{Worker: w},
		)
	}
}
