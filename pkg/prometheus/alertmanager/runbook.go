package alertmanager

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/rule_parser"

func (c *Client) Runbook(name string) string {
	return rule_parser.New(c.Rules()).Runbook(name)
}
