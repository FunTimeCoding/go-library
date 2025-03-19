package label_filter

import label2 "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/label_filter/label"

func (f *Filter) DropValue(
	label string,
	value string,
) {
	f.dropValue = append(f.dropValue, label2.New(label, value))
}
