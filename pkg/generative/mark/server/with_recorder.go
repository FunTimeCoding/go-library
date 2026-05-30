package server

import "github.com/funtimecoding/go-library/pkg/face"

func (b *Builder) WithRecorder(r face.Recorder) *Builder {
	b.recorder = r

	return b
}
