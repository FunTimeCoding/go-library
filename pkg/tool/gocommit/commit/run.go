package commit

import (
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/common"
	"github.com/funtimecoding/go-library/pkg/tool/gocommit/commit/option"
)

func Run(o *option.Commit) {
	validate(o)
	c := gitlab.New(o.Host, o.Token)
	project := common.FindProjectOrExit(c, o.Owner, o.Repository)
	c.Commit(
		project.Identifier,
		o.Branch,
		o.Message,
		o.Path,
		strings.ReplaceAllSlice(
			system.ReadFile(o.Template),
			o.Replace,
		),
		c.FileExists(project, o.Branch, o.Path),
	)
}
