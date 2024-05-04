package go_mod

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
	"os/exec"
)

func Update(
	name string,
	skipProxy bool,
) {
	fmt.Printf("%s skip=%v\n", name, skipProxy)

	if skipProxy {
		s := []string{"go", "get", name}
		c := exec.Command(s[0], s[1:]...)
		c.Env = os.Environ()
		c.Env = append(c.Env, "GOPROXY=direct")
		combined, e := c.CombinedOutput()

		if t := string(combined); t != "" {
			fmt.Printf("Output:\n%s", t)
		}

		errors.PanicOnError(e)
	} else {
		system.Run("go", "get", name)
	}
}
