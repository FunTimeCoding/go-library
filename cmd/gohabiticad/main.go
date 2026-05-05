package main

import "github.com/funtimecoding/go-library/pkg/tool/gohabiticad"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gohabiticad.Main(Version, GitHash, BuildDate)
}
