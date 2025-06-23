package file

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

func New(b *[]byte) *File {
	return &File{source: b}
}

func (f *File) Parse() {
	p := goldmark.DefaultParser()
	o := p.Parse(text.NewReader(*f.source))

	if false {
		Walk(f.source, o)
	}

	if true {
		WalkTree(f.source, o)
	}
}
