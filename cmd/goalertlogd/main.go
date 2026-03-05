package main

import "github.com/funtimecoding/go-library/pkg/tool/goalertlogd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goalertlogd.Main(Version, GitHash, BuildDate)
}
