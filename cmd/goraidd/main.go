package main

import "github.com/funtimecoding/go-library/pkg/tool/goraidd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goraidd.Main(Version, GitHash, BuildDate)
}
