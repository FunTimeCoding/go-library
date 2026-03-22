package check

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/linux"
	"github.com/funtimecoding/go-library/pkg/linux/systemd/command"
	"github.com/funtimecoding/go-library/pkg/linux/systemd/jc"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func Netstat(verbose bool) []*jc.Output {
	output := Execute(command.Netstat())

	if verbose {
		fmt.Printf("Netstat raw: %s\n", output)
	}

	var result []*jc.Output
	notation.DecodeStrict(
		Pipe(
			Pipe(output, verbose, linux.Awk, "!seen[$4]++"),
			verbose,
			linux.Jc, "--netstat", "--monochrome",
		),
		&result,
		true,
	)

	return result
}
