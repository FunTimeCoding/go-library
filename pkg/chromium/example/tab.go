package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chromium"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Tab() {
	c := chromium.NewEnvironment()
	defer c.Close()
	a := c.TabByHost(environment.Get("CHROMIUM_EXAMPLE_TAB", 1))

	if false {
		fmt.Printf("Tab: %s\n", a.Title)
		fmt.Printf("  ID %s\n", a.Id)
		fmt.Printf("  URL %s\n", a.Url)
		fmt.Printf("  DevtoolsFrontendUrl: %s\n", a.DevtoolsFrontendUrl)
		fmt.Printf("  WebSocketDebuggerUrl: %s\n", a.WebSocketDebuggerUrl)
	}

	fmt.Printf("Body: %+v", c.Body(a.Id))

	if false {
		for _, t := range chromium.NewEnvironment().Tabs() {
			fmt.Printf("Tab: %+v\n", t)
		}
	}
}
