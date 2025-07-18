package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
)

func Project() {
	for _, p := range gitlab.NewEnvironment().Projects() {
		fmt.Printf("Project: %s\n", p.Format(constant.Format))
	}
}
