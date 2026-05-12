package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/constant"
)

func BranchRequest() {
	a := argument.NewSimple("github-branch-request")
	a.String(argument.Branch, "", "Branch name")
	a.ParseSimple()
	branch := a.GetString(argument.Branch)
	c := github.NewEnvironment()
	f := option.ExtendedColor.Copy()
	fmt.Println(
		c.MustBranchRequest(
			constant.LibraryNamespace,
			constant.LibraryRepository,
			branch,
		).Format(f),
	)
}
