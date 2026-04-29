package convert

import (
	"github.com/funtimecoding/go-library/pkg/habitica"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/generated/server"
)

func ChecklistItems(v []habitica.ChecklistItem) []server.ChecklistItem {
	result := make([]server.ChecklistItem, 0, len(v))

	for _, i := range v {
		result = append(result, ChecklistItem(i))
	}

	return result
}
