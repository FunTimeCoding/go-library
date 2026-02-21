package main

import "github.com/funtimecoding/go-library/pkg/tool/goam"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goam.Main(Version, GitHash, BuildDate)
}
