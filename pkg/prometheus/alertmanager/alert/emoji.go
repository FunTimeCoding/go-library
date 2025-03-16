package alert

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"

func (a *Alert) emoji() []string {
	var result []string

	if a.Silenced() {
		result = append(result, "🔇")
	}

	if a.Documentation != constant.None {
		result = append(result, "🖊 ")
	}

	return result
}
