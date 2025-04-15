package hypertext

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func Title(d *goquery.Document) string {
	var result string
	d.Find("title").Each(
		func(
			index int,
			s *goquery.Selection,
		) {
			result = strings.TrimSpace(s.Text())
		},
	)

	return result
}
