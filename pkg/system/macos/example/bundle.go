package example

import "github.com/funtimecoding/go-library/pkg/system/macos"

func Bundle() {
	macos.CreateBundle(
		"Example",
		"tmp",
		"tmp/example",
		"tmp/icon.icns",
		"sh.s3n",
		"1.0.0",
	)
}
