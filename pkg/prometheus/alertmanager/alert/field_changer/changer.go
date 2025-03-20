package field_changer

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/field_changer/severity"

type Changer struct {
	name     map[string]string
	severity []*severity.Severity
}
