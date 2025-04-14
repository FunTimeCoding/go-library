package internal

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert/alert_enricher"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert/detail"
	opsgenieConstant "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/separator"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"strings"
)

func Opsgenie() *opsgenie.Client {
	result := opsgenie.NewEnvironment().Set(
		alert_enricher.New().Add(
			constant.HighMemoryUsage,
			constant.Memory,
			constant.High,
		),
	)

	if s := environment.GetDefault(
		opsgenieConstant.TeamEnvironment,
		"",
	); s != "" {
		for _, pair := range split.Semicolon(s) {
			if len(pair) == 0 {
				continue
			}

			k, v := key_value.Comma(pair)

			if len(k) == 0 || len(v) == 0 {
				continue
			}

			result.TeamMap().AddKey(k, v)
		}
	} else {
		result.TeamMap().AddKey("Infinite Loopsies", "INF")
	}

	result.ShortAlert(
		func(s string) string {
			if false {
				switch s {
				case constant.HighMemoryUsage:
					return "Memory"
				}
			}

			return s
		},
	)
	result.ShortUser(
		func(s string) string {
			if strings.Contains(s, separator.At) {
				k, _ := key_value.At(s)

				return k
			}

			return s
		},
	)
	result.DescriptionToName(
		func(s string) string {
			return s
		},
	)
	result.TagToTeam(
		func(s []string) string {
			return ""
		},
	)
	result.ParseDescription(
		func(s string) *detail.Prometheus {
			return detail.New()
		},
	)

	return result
}
