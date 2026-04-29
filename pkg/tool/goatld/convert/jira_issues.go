package convert

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/generated/server"
)

func JiraIssues(v []*issue.Issue) []*server.JiraIssue {
	result := make([]*server.JiraIssue, 0, len(v))

	for _, i := range v {
		result = append(result, JiraIssue(i))
	}

	return result
}
