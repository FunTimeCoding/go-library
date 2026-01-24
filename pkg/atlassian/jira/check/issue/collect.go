package issue

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/query"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/common"
)

func collect() []*issue.Issue {
	return query.Open(
		common.Jira(),
		environment.Required(constant.DefaultProjectNameEnvironment),
	)
}
