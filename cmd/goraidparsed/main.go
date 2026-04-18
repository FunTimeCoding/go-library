package main

import "github.com/funtimecoding/go-library/pkg/tool/goraidparsed"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goraidparsed.Main(Version, GitHash, BuildDate)
}
