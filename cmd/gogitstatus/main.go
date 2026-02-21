package main

import "github.com/funtimecoding/go-library/pkg/tool/gogitstatus"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gogitstatus.Main(Version, GitHash, BuildDate)
}
