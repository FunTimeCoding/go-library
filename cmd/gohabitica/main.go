package main

import "github.com/funtimecoding/go-library/pkg/tool/gohabitica"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gohabitica.Main(Version, GitHash, BuildDate)
}
