package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
)

func MergeRequest() {
	g := gitlab.NewEnvironment()
	f := constant.Format

	for _, r := range g.ProjectsMergeRequests() {
		fmt.Println(r.Format(f))
	}
}
