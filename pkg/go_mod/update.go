package go_mod

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
	"os/exec"
)

func Update(
	name string,
	skipProxy bool,
) {
	fmt.Printf("%s\n", name)

	if skipProxy {
		s := []string{constant.Go, constant.Get, name}
		c := exec.Command(s[0], s[1:]...)
		c.Env = os.Environ()
		c.Env = append(
			c.Env,
			key_value.Equals(constant.Proxy, constant.Direct),
		)
		combined, e := c.CombinedOutput()

		if t := string(combined); t != "" {
			strings.PrintTrim(t)
		}

		errors.PanicOnError(e)
	} else {
		strings.PrintTrim(system.Run(constant.Go, constant.Get, name))
	}
}
