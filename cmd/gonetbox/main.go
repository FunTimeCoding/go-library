package main

import "github.com/funtimecoding/go-library/pkg/tool/gonetbox"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gonetbox.Main(Version, GitHash, BuildDate)
}
