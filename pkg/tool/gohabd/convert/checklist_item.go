package convert

import (
	"github.com/funtimecoding/go-library/pkg/habitica"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/generated/server"
)

func ChecklistItem(i habitica.ChecklistItem) server.ChecklistItem {
	return server.ChecklistItem{
		Identifier: i.ID,
		Text:       i.Text,
		Completed:  i.Completed,
	}
}
