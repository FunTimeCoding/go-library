package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func BranchRequest() {
	pflag.Int64(argument.Project, 0, "Project ID")
	pflag.String(argument.Branch, "", "Branch name")
	argument.ParseBind()
	project := viper.GetInt64(argument.Project)
	branch := viper.GetString(argument.Branch)
	g := gitlab.NewEnvironment()
	f := constant.Format
	fmt.Println(g.BranchRequest(project, branch).Format(f))
}
