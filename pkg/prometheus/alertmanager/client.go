package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/prometheus"
	"github.com/funtimecoding/go-library/pkg/prometheus/rule/rule_list"
	"github.com/prometheus/alertmanager/api/v2/client"
)

type Client struct {
	client     *client.AlertmanagerAPI
	host       string
	prometheus *prometheus.Client
	rules      *rule_list.List
}
