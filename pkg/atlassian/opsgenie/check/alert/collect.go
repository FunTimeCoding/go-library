package alert

import (
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert"
)

func collect() []*alert.Alert {
	return internal.Opsgenie().Open()
}
