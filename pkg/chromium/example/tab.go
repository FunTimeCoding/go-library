package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chromium"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Tab() {
	c := chromium.NewEnvironment()
	defer c.Close()
	t := c.TabByHost(environment.Get("CHROMIUM_EXAMPLE_TAB"))

	if t == nil {
		panic("tab not found")
	}

	fmt.Printf("Body: %+v", c.Body(t.Id))
}
