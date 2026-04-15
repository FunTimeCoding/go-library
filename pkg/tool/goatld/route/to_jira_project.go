package route

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
)

func toJiraProjects(projects *jira.ProjectList) []server.JiraProject {
	result := make([]server.JiraProject, 0, len(*projects))

	for _, p := range *projects {
		entry := server.JiraProject{
			Key:  p.Key,
			Name: p.Name,
		}

		if p.ProjectTypeKey != "" {
			entry.Type = &p.ProjectTypeKey
		}

		if p.Self != "" {
			entry.Link = &p.Self
		}

		result = append(result, entry)
	}

	return result
}
