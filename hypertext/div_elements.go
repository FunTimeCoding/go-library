package hypertext

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func DivElements(d *goquery.Document) []string {
	var result []string
	d.Find("div").Each(
		func(
			index int,
			div *goquery.Selection,
		) {
			div.Find("dl").Each(
				func(
					indexDl int,
					dl *goquery.Selection,
				) {
					dl.Find("dt,dd").Each(
						func(
							indexDtDd int,
							dtDd *goquery.Selection,
						) {
							var text = strings.TrimSpace(dtDd.Text())

							if text != "" {
								result = append(result, text)
							}
						},
					)
				},
			)
		},
	)

	return result
}
