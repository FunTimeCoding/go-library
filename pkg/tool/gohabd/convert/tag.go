package convert

import (
	"github.com/funtimecoding/go-library/pkg/habitica"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/generated/server"
)

func Tag(t habitica.Tag) server.Tag {
	return server.Tag{
		Identifier: t.ID,
		Name:       t.Name,
	}
}
