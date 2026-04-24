package convert

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/habitica"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/server"
)

func ChecklistItem(i habitica.ChecklistItem) server.ChecklistItem {
	return server.ChecklistItem{
		Identifier: i.ID,
		Text:       i.Text,
		Completed:  i.Completed,
	}
}
