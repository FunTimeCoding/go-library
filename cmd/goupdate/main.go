package main

import "github.com/funtimecoding/go-library/pkg/tool/goupdate"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goupdate.Main(Version, GitHash, BuildDate)
}
