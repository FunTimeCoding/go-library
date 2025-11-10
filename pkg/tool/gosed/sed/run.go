package sed

import (
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/strings/base64"
	"github.com/funtimecoding/go-library/pkg/tool/common"
	"github.com/funtimecoding/go-library/pkg/tool/gosed/sed/option"
	"log"
)

func Run(o *option.Sed) {
	validate(o)
	c := gitlab.New(o.Host, o.Token)
	p := common.FindProjectOrExit(c, o.Owner, o.Repository)
	f := c.File(p.Identifier, o.Branch, o.Path)

	if f == nil {
		log.Panicf("file does not exist: %s", o.Path)
	}

	c.Commit(
		p.Identifier,
		o.Branch,
		o.Message,
		o.Path,
		replaceLines(base64.Decode(f.Content), o.Replaces),
		true,
	)
}
