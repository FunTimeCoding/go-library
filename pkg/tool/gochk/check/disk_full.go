package check

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	stringLibrary "github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"os/exec"
	"strings"
)

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
