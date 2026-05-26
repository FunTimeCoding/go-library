package main

import "github.com/funtimecoding/go-library/pkg/tool/gochromed"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gochromed.Main(Version, GitHash, BuildDate)
}
