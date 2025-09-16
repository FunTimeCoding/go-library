package issue

import (
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/query"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func collect() []*issue.Issue {
	return query.Open(
		internal.Jira(),
		environment.Get(constant.DefaultProjectNameEnvironment),
	)
}
