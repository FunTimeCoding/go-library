package task

import "github.com/xanzy/go-gitlab"

type Task struct {
	Body  string
	State string
	Link  string
	Type  gitlab.TodoTargetType

	Raw *gitlab.Todo
}
