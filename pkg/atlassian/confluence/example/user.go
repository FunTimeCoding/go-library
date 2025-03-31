package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func User() {
	c := confluence.NewEnvironment()
	f := constant.DefaultFormat

	if true {
		fmt.Println(c.User().Format(f))
	}

	if false {
		fmt.Println("UserVirtomize")
		fmt.Println(c.UserVirtomize())
	}

	if false {
		fmt.Println("UserKaos")
		fmt.Println(c.UserKaos())
	}
}
