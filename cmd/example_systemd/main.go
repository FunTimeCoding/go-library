package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/linux/systemd"
	"github.com/funtimecoding/go-library/pkg/ssh"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/time"
	"github.com/spf13/pflag"
	"os"
)

func main() {
	pflag.Parse()
	host := pflag.Arg(0)

	if host == "" {
		fmt.Println("Host required")

		os.Exit(1)
	}

	service := pflag.Arg(1)

	if service == "" {
		fmt.Println("Service required")

		os.Exit(1)
	}

	fmt.Printf("Host: %s\n", host)
	s := ssh.New(system.User().Username, host, true)
	defer s.Close()
	c := systemd.New(s)

	if false {
		fmt.Printf("List: %+v\n", c.ListOutput())
		fmt.Printf("Not found: %+v\n", c.NotFoundOutput())
		fmt.Printf("Status: %+v\n", c.Status(service))
	}

	if true {
		fmt.Printf(
			"Is active running: %v\n",
			c.IsActiveRunning(service),
		)
		fmt.Printf(
			"Start time: %v\n",
			time.Format(c.StartTime(service)),
		)
	}
}
