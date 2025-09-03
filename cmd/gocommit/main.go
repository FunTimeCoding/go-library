package main

import "github.com/funtimecoding/go-library/pkg/tool/gocommit"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gocommit.Main(Version, GitHash, BuildDate)
}
