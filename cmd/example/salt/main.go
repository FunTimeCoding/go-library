package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/provision/salt"
	"github.com/funtimecoding/go-library/pkg/provision/salt/constant"
)

func main() {
	c := salt.NewEnvironment()

	if false {
		fmt.Println(c.Minions())
	}

	if true {
		fmt.Println(
			c.Local("saltstack-test", constant.Run, []string{"whoami"}),
		)
	}
}
