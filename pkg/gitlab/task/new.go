package task

import "github.com/xanzy/go-gitlab"

func New(v *gitlab.Todo) *Task {
	return &Task{
		Type:  v.TargetType,
		Body:  v.Body,
		State: v.State,
		Link:  v.TargetURL,
		Raw:   v,
	}
}
