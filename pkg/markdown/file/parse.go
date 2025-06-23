package file

import (
	"github.com/funtimecoding/go-library/pkg/markdown/file/flat"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

func (f *File) Parse() *flat.Flat {
	p := goldmark.DefaultParser()
	o := p.Parse(text.NewReader(*f.source))
	l := flat.New()

	if false {
		Walk(f.source, o, l)
	}

	if true {
		WalkTree(f.source, o, l)
	}

	return l
}
