package project

import "github.com/funtimecoding/go-library/pkg/console/description"

func (p *Project) Meta() *description.Description {
	return description.New("Project", "Project")
}
