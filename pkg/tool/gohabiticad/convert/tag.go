package convert

import (
	"github.com/funtimecoding/go-library/pkg/habitica/tag"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
)

func Tag(t *tag.Tag) *server.Tag {
	return &server.Tag{Identifier: t.ID, Name: t.Name}
}
