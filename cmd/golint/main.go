package main

import "github.com/funtimecoding/go-library/pkg/tool/golint"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	golint.Main(Version, GitHash, BuildDate)
}
