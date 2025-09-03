package go_mod

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/system/run"
)

func Update(
	name string,
	skipProxy bool,
	continueOnError bool,
) {
	fmt.Printf("%s\n", name)
	r := run.New()
	r.Panic = false

	if skipProxy {
		r.Environment(constant.Proxy, constant.Direct)
	}

	r.Start(constant.Go, constant.Get, name)
	r.Print()

	if !continueOnError {
		r.PanicOnError()
	}
}
