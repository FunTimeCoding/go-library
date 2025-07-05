package git

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/git/repository"
	"os/exec"
)

func ReadRepository(path string) *repository.Repository {
	c := exec.Command(constant.Command, constant.RevParse, constant.GitDir)
	c.Dir = path
	errors.PanicOnError(c.Run())
	c = exec.Command(constant.Command, constant.Status, constant.Porcelain)
	c.Dir = path
	output, e := c.Output()
	errors.PanicOnError(e)

	if false {
		fmt.Printf("Output: %s\n", output)
	}

	return repository.New(path, string(output))
}
