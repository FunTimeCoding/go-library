package check

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/linux"
	"github.com/funtimecoding/go-library/pkg/linux/systemd/command"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"github.com/spf13/viper"
	"runtime"
	"slices"
)

func Run() {
	switch runtime.GOOS {
	case constant.Linux:
		fmt.Println("Linux")
		fmt.Printf("Cores: %d\n", runtime.NumCPU())
		fmt.Printf("Failed: %s\n", RunCommand(command.Failed()))
		// TODO: Load average > CPU cores check
		diskFull()
		if run.CommandExists(linux.Jc) {
			if port := viper.GetString(argument.Port); port != "" {
				ports := split.Comma(port)
				var found []string

				for _, n := range Netstat(false) {
					if n.NetworkProtocol != "ipv4" {
						continue
					}

					if !slices.Contains(ports, n.LocalPort) {
						continue
					}

					found = append(found, n.LocalPort)
					fmt.Printf("Found port: %s\n", n.LocalPort)
				}

				slices.Sort(ports)
				slices.Sort(found)

				if !slices.Equal(ports, found) {
					fmt.Printf(
						"Expect ports: %s\n",
						join.Comma(ports),
					)
					fmt.Printf(
						"Found ports: %s\n",
						join.Comma(found),
					)
				}
			}
		} else {
			fmt.Println("jc not found")
		}
	case constant.Darwin:
		fmt.Println("Darwin")
	}
}
