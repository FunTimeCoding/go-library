package convert

import (
	"github.com/funtimecoding/go-library/pkg/habitica/response"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
)

func Tasks(v []response.Task) []server.Task {
	result := make([]server.Task, 0, len(v))

	for _, t := range v {
		result = append(result, Task(t))
	}

	return result
}
