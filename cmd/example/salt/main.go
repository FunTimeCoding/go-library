package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/provision/salt"
	"github.com/funtimecoding/go-library/pkg/provision/salt/constant"
)

func main() {
	c := salt.NewEnvironment()

	if false {
		for _, m := range c.MustMinions() {
			fmt.Printf("Minion: %+v\n", m)
		}
	}

	if false {
		for k, v := range c.MustLocal(
			"saltstack-test",
			constant.Run,
			[]string{"whoami"},
		) {
			fmt.Printf("%s: %+v\n", k, v)
		}
	}

	if false {
		identifier := c.MustLocalAsync(
			"saltstack-test",
			constant.Run,
			[]string{"sleep 30 && echo hello"},
		)
		fmt.Printf("JID: %s\n", identifier)
		fmt.Printf("Job: %+v", c.MustJob(identifier))
	}

	if true {
		for _, j := range c.MustJobs() {
			fmt.Printf("Job: %+v\n", j)
		}
	}
}
