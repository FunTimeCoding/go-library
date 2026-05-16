package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/tag"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
)

func Tag(t *tag.Tag) *server.Tag {
	return &server.Tag{Identifier: t.Identifier, Name: t.Name}
}
