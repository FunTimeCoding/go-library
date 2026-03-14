package example

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/funtimecoding/go-library/pkg/system/macos"
)

func Bundle() {
	macos.CreateBundle(
		"Example",
		"tmp",
		"tmp/example",
		"tmp/icon.icns",
		"sh.s3n",
		constant.DefaultVersion,
	)
}
