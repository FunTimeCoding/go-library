package main

import "github.com/funtimecoding/go-library/pkg/tool/goprometheusd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goprometheusd.Main(Version, GitHash, BuildDate)
}
