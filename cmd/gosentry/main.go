package main

import "github.com/funtimecoding/go-library/pkg/tool/gosentry"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosentry.Main(Version, GitHash, BuildDate)
}
