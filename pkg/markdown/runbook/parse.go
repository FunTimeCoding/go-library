package runbook

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

func (r *Runbook) Parse(filename string) {
	o := goldmark.DefaultParser().Parse(text.NewReader(*r.source))
	r.Filename = filename
	r.Walk(o)
}
