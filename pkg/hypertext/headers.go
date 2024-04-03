package hypertext

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func Headers(d *goquery.Document) (result []string) {
	d.Find("h1").Each(
		func(
			_ int,
			s *goquery.Selection,
		) {
			if t := strings.TrimSpace(s.Text()); t != "" {
				result = append(result, t)
			}
		},
	)
	d.Find("h2").Each(
		func(
			_ int,
			s *goquery.Selection,
		) {
			if t := strings.TrimSpace(s.Text()); t != "" {
				result = append(result, t)
			}
		},
	)
	d.Find("h3").Each(
		func(
			_ int,
			s *goquery.Selection,
		) {
			if t := strings.TrimSpace(s.Text()); t != "" {
				result = append(result, t)
			}
		},
	)
	d.Find("h4").Each(
		func(
			_ int,
			s *goquery.Selection,
		) {
			if t := strings.TrimSpace(s.Text()); t != "" {
				result = append(result, t)
			}
		},
	)
	d.Find("h5").Each(
		func(
			_ int,
			s *goquery.Selection,
		) {
			if t := strings.TrimSpace(s.Text()); t != "" {
				result = append(result, t)
			}
		},
	)
	d.Find("h6").Each(
		func(
			_ int,
			s *goquery.Selection,
		) {
			if t := strings.TrimSpace(s.Text()); t != "" {
				result = append(result, t)
			}
		},
	)

	return
}
