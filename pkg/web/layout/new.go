package layout

import "github.com/funtimecoding/go-library/pkg/identity"

func New(i *identity.Tool) *Page {
	return &Page{identity: i}
}
