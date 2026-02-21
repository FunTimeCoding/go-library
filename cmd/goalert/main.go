package main

import "github.com/funtimecoding/go-library/pkg/tool/goalert"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goalert.Main(Version, GitHash, BuildDate)
}
