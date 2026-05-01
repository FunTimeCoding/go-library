package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func User() {
	fmt.Println(confluence.NewEnvironment().MustUser().Format(constant.Format))
}
