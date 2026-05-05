package main

import "github.com/funtimecoding/go-library/pkg/tool/gopostgres"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gopostgres.Main(Version, GitHash, BuildDate)
}
