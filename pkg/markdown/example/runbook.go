package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/markdown"
	"github.com/funtimecoding/go-library/pkg/markdown/constant"
	"github.com/funtimecoding/go-library/pkg/markdown/runbook"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Runbook() {
	base := environment.Required(constant.WikiPathEnvironment)
	f := option.Color

	for _, n := range system.Files(base) {
		fmt.Printf("File: %s\n", n)
		source := system.ReadBytes(base, n)

		if false {
			markdown.Print(source, f)
		}

		r := runbook.New(&source)
		r.Parse(n)
		fmt.Printf("Runbook: %+v\n", r)
	}
}
