package main

import "github.com/funtimecoding/go-library/pkg/tool/goterraformd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goterraformd.Main(Version, GitHash, BuildDate)
}
