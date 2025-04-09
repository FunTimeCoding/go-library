package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/prometheus"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/alert_enricher"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/field_changer"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/label_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/name_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/rule/rule_list"
	"github.com/prometheus/alertmanager/api/v2/client"
)

type Client struct {
	client      *client.AlertmanagerAPI
	host        string
	prometheus  *prometheus.Client
	rules       *rule_list.List
	enricher    *alert_enricher.Enricher
	changer     *field_changer.Changer
	nameFilter  *name_filter.Filter
	labelFilter *label_filter.Filter
}
