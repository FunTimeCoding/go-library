package main

import "github.com/funtimecoding/go-library/pkg/tool/goanalyze"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goanalyze.Main(Version, GitHash, BuildDate)
}
