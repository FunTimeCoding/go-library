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
	pflag.String(argument.Host, "", "Relay host")
	pflag.String(TargetHost, "", "Target host")
	pflag.Int(TargetPort, 0, "Target port")
	argument.ParseBind()
	t := tunnel.New()

	if false {
		t.Verbose = true
		t.NoOutput = true
	}

	fmt.Printf("Start: %+v\n", t)
	t.Start(
		argument.Required(argument.Host),
		argument.Required(TargetHost),
		argument.RequiredInteger(TargetPort),
		0,
	)
	defer t.Stop()
	fmt.Println("Sleep 10")
	system.Sleep(10)
	fmt.Println("Stop")
}
