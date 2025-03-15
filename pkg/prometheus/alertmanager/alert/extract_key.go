package alert

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func extractKey(
	remaining *models.LabelSet,
	k string,
	to *string,
) {
	if v, ok := (*remaining)[k]; ok {
		*to = v
		delete(*remaining, k)
	} else {
		*to = constant.None
	}
}
