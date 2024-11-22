package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/ssh/tunnel"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/spf13/pflag"
)

const (
	TargetHost = "target-host"
	TargetPort = "target-port"
)

func main() {
	pflag.StringP(argument.Host, "o", "", "Relay host")
	pflag.StringP(TargetHost, "t", "", "Target host")
	pflag.IntP(TargetPort, "p", 0, "Target port")
	argument.ParseAndBind()
	t := tunnel.New()

	if false {
		t.Verbose = true
		t.NoOutput = true
	}

	fmt.Printf("Start: %+v\n", t)
	t.Start(
		argument.RequiredStringFlag(argument.Host, 1),
		argument.RequiredStringFlag(TargetHost, 1),
		argument.RequiredIntegerFlag(TargetPort, 1),
		0,
	)
	defer t.Stop()
	fmt.Println("Sleep 10")
	system.Sleep(10)
	fmt.Println("Stop")
}
