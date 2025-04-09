package internal

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/alert_enricher"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/field_changer"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/label_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/name_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
)

func Alertmanager() *alertmanager.Client {
	return alertmanager.NewEnvironment().Set(
		alert_enricher.New().Add(
			constant.KubernetesCronJobFailed,
			constant.Job,
			constant.Fail,
		),
		field_changer.New(),
		name_filter.New(true),
		label_filter.New(true),
	)
}
