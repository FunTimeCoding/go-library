package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	libraryProject "github.com/funtimecoding/go-library/pkg/project"
)

func (c *Client) PipelineProjects() []*project.Project {
	return c.ProjectsWithFile(libraryProject.GitLabFile)
}
