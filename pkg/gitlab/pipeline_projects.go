package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	library "github.com/funtimecoding/go-library/pkg/project"
)

func (c *Client) PipelineProjects() []*project.Project {
	return c.ProjectsWithFile(library.GitLabFile, false)
}
