package example

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/go-library/pkg/chromium"
	"log"
	"strings"
	"time"
)

func Open() {
	c := chromium.NewEnvironment()
	defer c.Close()
	var result string
	c.Run(
		chromedp.Navigate(`https://pkg.go.dev/time`),
		//chromedp.Text(
		//	`.Documentation-overview`,
		//	&result,
		//	chromedp.NodeVisible,
		//),
		chromedp.Sleep(1*time.Second),
	)
	log.Println(strings.TrimSpace(result))
}
