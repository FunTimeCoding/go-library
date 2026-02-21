package main

import "github.com/funtimecoding/go-library/pkg/tool/gomonitor"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gomonitor.Main(Version, GitHash, BuildDate)
}
