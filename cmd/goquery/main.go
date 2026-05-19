package main

import "github.com/funtimecoding/go-library/pkg/tool/goquery"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goquery.Main(Version, GitHash, BuildDate)
}
