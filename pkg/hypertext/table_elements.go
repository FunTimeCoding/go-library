package hypertext

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func TableElements(d *goquery.Document) []string {
	var result []string
	d.Find("table").Each(
		func(
			index int,
			t *goquery.Selection,
		) {
			t.Find("tr").Each(
				func(
					indexRow int,
					tr *goquery.Selection,
				) {
					tr.Find("td").Each(
						func(
							indexValue int,
							td *goquery.Selection,
						) {
							var text = strings.TrimSpace(td.Text())

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
