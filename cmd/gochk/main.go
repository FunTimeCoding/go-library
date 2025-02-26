package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/linux/systemd/command"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"runtime"
	"strings"
)

func main() {
	switch runtime.GOOS {
	case constant.Linux:
		fmt.Println("Linux")
		fmt.Printf("Cores: %d\n", runtime.NumCPU())
		fmt.Printf("Failed: %+v\n", Run(command.Failed()))
		// TODO: Load average > CPU cores check
		// TODO: Disk full check
		// TODO: Expected port
		//  443 can be found in ip -all netns exec netstat -plnt
	case constant.Darwin:
		fmt.Println("Darwin")
	}
}

func Run(c []string) string {
	return strings.TrimSpace(run.New().Start(c...))
}
