package label_filter

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/label_filter/label"

type Filter struct {
	keep          []string
	drop          []string
	keepValue     []*label.Label
	dropValue     []*label.Label
	KeepByDefault bool
}
