package main

import "github.com/funtimecoding/go-library/pkg/tool/goalertlog"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goalertlog.Main(Version, GitHash, BuildDate)
}
