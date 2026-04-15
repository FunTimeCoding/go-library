package route

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
)

func toJiraIssues(issues []*issue.Issue) []server.JiraIssue {
	result := make([]server.JiraIssue, 0, len(issues))

	for _, i := range issues {
		result = append(result, toJiraIssue(i))
	}

	return result
}
