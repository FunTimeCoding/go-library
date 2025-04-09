package internal

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert/alert_enricher"
)

func Opsgenie() *opsgenie.Client {
	return opsgenie.NewEnvironment().Set(alert_enricher.New())
}
