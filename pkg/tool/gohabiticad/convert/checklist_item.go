package convert

import (
	"github.com/funtimecoding/go-library/pkg/habitica/checklist_item"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
)

func ChecklistItem(i *checklist_item.Item) *server.ChecklistItem {
	return &server.ChecklistItem{
		Identifier: i.ID,
		Text:       i.Text,
		Completed:  i.Completed,
	}
}
