package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func BranchRequest() {
	pflag.String(argument.Branch, "", "Branch name")
	argument.ParseBind()
	branch := viper.GetString(argument.Branch)
	c := github.NewEnvironment()
	f := option.ExtendedColor.Copy()
	fmt.Println(
		c.BranchRequest(
			constant.LibraryNamespace,
			constant.LibraryRepository,
			branch,
		).Format(f),
	)
}
