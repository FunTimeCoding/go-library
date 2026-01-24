package alert

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert"
	"github.com/funtimecoding/go-library/pkg/tool/common"
)

func collect() []*alert.Alert {
	return common.Opsgenie().Open()
}
