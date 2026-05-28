package main

import "github.com/funtimecoding/go-library/pkg/tool/goprocessd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goprocessd.Main(Version, GitHash, BuildDate)
}
