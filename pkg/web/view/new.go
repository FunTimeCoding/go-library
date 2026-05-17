package view

import "github.com/funtimecoding/go-library/pkg/web/layout"

func New(l *layout.Page) *View {
	return &View{layout: l}
}
