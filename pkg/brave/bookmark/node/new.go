package node

import "github.com/funtimecoding/go-library/pkg/brave/bookmark/file"

func New(v *file.Node) *Node {
	result := Stub()
	result.Type = v.Type
	result.Name = v.Name
	result.Link = v.Url
	result.Children = NewSlice(v.Children)
	result.Raw = v

	return result
}
