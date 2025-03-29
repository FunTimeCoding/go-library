package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/linux/systemd"
	"github.com/funtimecoding/go-library/pkg/ssh"
	"github.com/funtimecoding/go-library/pkg/ssh/command"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/time"
	"github.com/spf13/pflag"
)

func main() {
	pflag.Parse()
	host := argument.RequiredPositional(0, "HOST", 1)
	service := argument.RequiredPositional(
		1,
		"SERVICE",
		1,
	)
	fmt.Printf("Host: %s\n", host)
	fmt.Printf("Service: %s\n", service)

	s := ssh.New(system.User().Username, host, true)
	defer s.Close()
	c := systemd.New(s)

	if false {
		fmt.Printf("List: %+v\n", c.ListOutput())
		fmt.Printf("Not found: %+v\n", c.NotFoundOutput())
		fmt.Printf("Status: %+v\n", c.Status(service))
	}

	if false {
		fmt.Printf(
			"Is active running: %v\n",
			c.IsActiveRunning(service),
		)
		fmt.Printf(
			"Start time: %v\n",
			time.Format(c.StartTime(service)),
		)
	}

	if true {
		s.RunCommand(
			command.New("sudo echo $TEST").Set("TEST", "yay"),
		).Print()
	}
}
