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
	p := common.FindProjectOrExit(c, o.Owner, o.Repository)
	c.Commit(
		p.Identifier,
		o.Branch,
		o.Message,
		o.Path,
		strings.ReplaceAllSlice(
			system.ReadFileUnsafe(o.Template),
			o.Replace,
		),
		c.FileExists(p, o.Branch, o.Path),
	)
}
