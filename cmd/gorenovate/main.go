package main

import "github.com/funtimecoding/go-library/pkg/tool/gorenovate"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gorenovate.Main(Version, GitHash, BuildDate)
}
