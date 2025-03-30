package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
)

func User() {
	c := confluence.NewEnvironment()

	if true {
		fmt.Println("UserVirtomize")
		fmt.Println(c.UserVirtomize())
	}

	if true {
		fmt.Println("UserKaos")
		fmt.Println(c.UserKaos())
	}
}
