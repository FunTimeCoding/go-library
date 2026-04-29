package convert

import (
	"github.com/funtimecoding/go-library/pkg/habitica"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/generated/server"
)

func Tasks(v []habitica.Task) []server.Task {
	result := make([]server.Task, 0, len(v))

	for _, t := range v {
		result = append(result, Task(t))
	}

	return result
}
