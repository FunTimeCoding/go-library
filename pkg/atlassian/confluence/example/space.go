package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func Space() {
	f := constant.Format

	for _, s := range confluence.NewEnvironment().MustSpaces() {
		fmt.Println(s.Format(f))
	}
}
