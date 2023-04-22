package hypertext

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func Title(document *goquery.Document) string {
	var result string
	document.Find("title").Each(
		func(
			index int,
			element *goquery.Selection,
		) {
			result = strings.TrimSpace(element.Text())
		},
	)

	return result
}
