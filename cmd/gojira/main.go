package main

import "github.com/funtimecoding/go-library/pkg/tool/gojira"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gojira.Main(Version, GitHash, BuildDate)
}
