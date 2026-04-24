package convert

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/habitica"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/server"
)

func ChecklistItems(v []habitica.ChecklistItem) []server.ChecklistItem {
	result := make([]server.ChecklistItem, 0, len(v))

	for _, i := range v {
		result = append(result, ChecklistItem(i))
	}

	return result
}
