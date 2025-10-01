package sed

import (
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/tool/common"
	"github.com/funtimecoding/go-library/pkg/tool/gosed/sed/option"
)

func Run(o *option.Sed) {
	validate(o)
	c := gitlab.New(o.Host, o.Token)
	p := common.FindProjectOrExit(c, o.Owner, o.Repository)
	f := c.File(p.Identifier, o.Branch, o.Path)

	if f == nil {
		panic("file does not exist: " + o.Path)
	}

	c.Commit(
		p.Identifier,
		o.Branch,
		o.Message,
		o.Path,
		replaceLines(f.Content, o.Replaces),
		true,
	)
}
