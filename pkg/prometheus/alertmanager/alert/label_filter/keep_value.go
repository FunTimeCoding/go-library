package label_filter

import label2 "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/label_filter/label"

func (f *Filter) KeepValue(
	label string,
	value string,
) {
	f.keepValue = append(f.keepValue, label2.New(label, value))
}
