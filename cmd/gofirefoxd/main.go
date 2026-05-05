// Start: Load js/firefox_bridge as temporary add-on via about:debugging in Firefox
package main

import "github.com/funtimecoding/go-library/pkg/tool/gofirefoxd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gofirefoxd.Main(Version, GitHash, BuildDate)
}
