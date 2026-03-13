package task

import "gitlab.com/gitlab-org/api/client-go/v2"

type Task struct {
	Body  string
	State string
	Link  string
	Type  gitlab.TodoTargetType

	Raw *gitlab.Todo
}
