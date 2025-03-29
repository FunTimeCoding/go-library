package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/linux"
	"github.com/funtimecoding/go-library/pkg/linux/systemd/command"
	"github.com/funtimecoding/go-library/pkg/linux/systemd/jc"
	"github.com/funtimecoding/go-library/pkg/notation"
	stringLibrary "github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os/exec"
	"runtime"
	"slices"
	"strings"
)

func main() {
	pflag.String(
		argument.Port,
		"",
		"Port, multiple values separated by comma",
	)
	argument.ParseAndBind()

	switch runtime.GOOS {
	case constant.Linux:
		fmt.Println("Linux")
		fmt.Printf("Cores: %d\n", runtime.NumCPU())
		fmt.Printf("Failed: %s\n", Run(command.Failed()))
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

func Run(c []string) string {
	return strings.TrimSpace(run.New().Start(c...))
}

func Netstat(verbose bool) []*jc.Output {
	output := Run(command.Netstat())

	if verbose {
		fmt.Printf("Netstat raw: %s\n", output)
	}

	var result []*jc.Output
	notation.DecodeStrict(
		Pipe(
			exec.Command(linux.Jc, "--netstat", "--monochrome"),
			Pipe(
				exec.Command(linux.Awk, "!seen[$4]++"), output,
				verbose,
			),
			verbose,
		),
		&result,
		true,
	)

	return result
}

func Pipe(
	c *exec.Cmd,
	input string,
	verbose bool,
) string {
	stdout, stderr := system.Pipe(c, input)

	if verbose {
		if stdout != "" {
			fmt.Printf("Pipe output: %s\n", stdout)
		}

		if stderr != "" {
			fmt.Printf("Pipe error: %s\n", stderr)
		}
	}

	return stdout
}

func diskFull() {
	output, e := exec.Command("df", "-h").CombinedOutput()
	errors.PanicOnError(e)

	for _, line := range split.NewLine(string(output))[1:] {
		p := strings.Fields(line)

		if len(p) < 4 {
			continue
		}

		u := stringLibrary.ToIntegerStrict(
			strings.TrimSuffix(p[len(p)-2], "%"),
		)
		point := p[0]

		if u == 100 {
			fmt.Printf("Disk full: %s", point)
		} else if u > 90 {
			fmt.Printf("Disk almost full: %s", point)
		}
	}
}
