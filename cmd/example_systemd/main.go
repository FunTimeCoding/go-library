package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/linux/systemd"
	"github.com/funtimecoding/go-library/pkg/ssh"
	"github.com/funtimecoding/go-library/pkg/system"
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

	fmt.Printf("Host: %s\n", host)
	s := ssh.New(system.User().Username, host, true)
	defer s.Close()
	c := systemd.New(s)

	if false {
		fmt.Printf("List: %+v\n", c.List().OutputString)
	}

	if false {
		fmt.Printf("Not found: %+v\n", c.NotFound().OutputString)
	}

	if true {
		fmt.Printf(
			"Status: %+v\n",
			systemd.ParseStatus(
				systemd.GrabUntilFirstBlankLine(
					c.Status("ssh.service").OutputString,
				),
			),
		)
	}
}
