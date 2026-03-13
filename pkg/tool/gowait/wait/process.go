package wait

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gowait/wait/option"
	"os/exec"
	"time"
)

func Process(o *option.Wait) {
	start := time.Now()
	attempt := 0

	for {
		attempt++

		if o.Verbose {
			fmt.Printf("Check %d\n", attempt)
		}

		c := exec.Command("pgrep", "-f", o.Process)

		if e := c.Run(); e != nil {
			fmt.Printf(
				"Done after %v\n",
				time.Since(start).Round(time.Second),
			)

			return
		}

		if o.Verbose {
			verboseCommand := exec.Command(
				"pgrep",
				"-fa",
				o.Process,
			)
			out, f := verboseCommand.Output()
			errors.PanicOnError(f)
			fmt.Printf("Running: %s", string(out))
		}

		if time.Since(start) > o.Timeout {
			panic(fmt.Sprintf("timeout after %v", o.Timeout))
		}

		time.Sleep(interval)
	}
}
