package example

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	system "github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/macos"
)

func Bundle() {
	macos.CreateBundle(
		"Example",
		system.Temporary,
		"tmp/example",
		"tmp/icon.icns",
		"sh.s3n",
		constant.DefaultVersion,
	)
}
