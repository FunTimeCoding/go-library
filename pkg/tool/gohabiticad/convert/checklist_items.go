package convert

import (
	"github.com/funtimecoding/go-library/pkg/habitica/response"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
)

func ChecklistItems(v []response.ChecklistItem) []server.ChecklistItem {
	result := make([]server.ChecklistItem, 0, len(v))

	for _, i := range v {
		result = append(result, ChecklistItem(i))
	}

	return result
}
