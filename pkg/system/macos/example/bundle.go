package example

import (
	library "github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/macos"
)

func Bundle() {
	macos.CreateBundle(
		"Example",
		constant.Temporary,
		"tmp/example",
		"tmp/icon.icns",
		"sh.s3n",
		library.DefaultVersion,
	)
}
