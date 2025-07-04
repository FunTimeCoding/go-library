package status

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/git/repository"
	"os/exec"
)

func checkRepository(path string) *repository.Repository {
	c := exec.Command("git", "rev-parse", "--git-dir")
	c.Dir = path
	errors.PanicOnError(c.Run())
	c = exec.Command("git", "status", "--porcelain")
	c.Dir = path
	output, e := c.Output()
	fmt.Printf("Output: %s\n", output)
	errors.PanicOnError(e)

	return repository.New(path, string(output))
}
