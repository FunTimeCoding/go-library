package main

import "github.com/funtimecoding/go-library/pkg/tool/goansible"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goansible.Main(Version, GitHash, BuildDate)
}
