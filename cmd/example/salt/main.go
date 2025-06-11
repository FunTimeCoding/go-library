package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/provision/salt"
	"github.com/funtimecoding/go-library/pkg/provision/salt/constant"
)

func main() {
	c := salt.NewEnvironment()
	fmt.Println(c.Minions())
	fmt.Println(
		c.Local("minion1", constant.Run, []string{"whoami"}),
	)
}
