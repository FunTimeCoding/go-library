package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/markdown"
	"github.com/funtimecoding/go-library/pkg/markdown/constant"
	"github.com/funtimecoding/go-library/pkg/markdown/file"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func File() {
	base := environment.Get(constant.WikiPathEnvironment)
	f := option.Color

	for _, n := range system.Files(base) {
		fmt.Printf("File: %s\n", n)
		source := system.ReadBytes(system.Join(base, n))
		markdown.Print(source, f)
		i := file.New(&source)
		i.Parse()

		if true {
			break
		}
	}
}
