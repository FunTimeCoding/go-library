package task

import "github.com/funtimecoding/go-library/pkg/console/description"

func (t *Task) Meta() *description.Description {
	return description.New("Task", "Task")
}
