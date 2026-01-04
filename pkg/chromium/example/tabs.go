package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chromium"
	"github.com/funtimecoding/go-library/pkg/chromium/constant"
	library "github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
	"github.com/funtimecoding/go-library/pkg/system"
	systemConstant "github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"strings"
)

func Tabs() {
	c := chromium.NewEnvironment()

	if false {
		// Do not close tabs
		defer c.Close()
	}

	for _, t := range c.Tabs() {
		if t.Type != constant.PageTabType {
			continue
		}

		if !strings.HasSuffix(t.Url, library.GraphicExtension) {
			continue
		}

		name, _ := key_value.Space(t.Title)
		fmt.Printf("%s %s\n", name, t.Url)
		p := join.Absolute(system.Home(), systemConstant.DownloadsPath, name)

		if !system.FileExists(p) {
			if c.NeedReload(t.Id, t.Url) {
				c.Activate(t.Id)
			}

			c.Save(c.TargetContext(t.Id), t.Url, p)
		} else {
			fmt.Println("  Exists")
		}
	}
}
