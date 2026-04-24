package convert

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/habitica"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/server"
)

func Tasks(v []habitica.Task) []server.Task {
	result := make([]server.Task, 0, len(v))

	for _, t := range v {
		result = append(result, Task(t))
	}

	return result
}
