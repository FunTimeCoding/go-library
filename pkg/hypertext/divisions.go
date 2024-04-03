package hypertext

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func Divisions(d *goquery.Document) (result []string) {
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
							if s := strings.TrimSpace(dtDd.Text()); s != "" {
								result = append(result, s)
							}
						},
					)
				},
			)
		},
	)

	return
}
