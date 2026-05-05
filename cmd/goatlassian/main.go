package main

import "github.com/funtimecoding/go-library/pkg/tool/goatlassian"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goatlassian.Main(Version, GitHash, BuildDate)
}
