package main

import "github.com/funtimecoding/go-library/pkg/tool/gokubernetesd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gokubernetesd.Main(Version, GitHash, BuildDate)
}
