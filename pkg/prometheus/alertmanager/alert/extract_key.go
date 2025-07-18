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
	if v, okay := (*remaining)[k]; okay {
		*to = v
		delete(*remaining, k)
	} else {
		*to = constant.None
	}
}
