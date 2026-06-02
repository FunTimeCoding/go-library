package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/go-library/pkg/tool/common"
)

func Alert() {
	alerts, _ := common.Alertmanager().MustAlerts(advanced_option.New(), nil)
	f := option.ExtendedColor.Copy()

	for _, a := range alerts {
		fmt.Printf("Alert: %+v\n", a.Format(f))
	}
}
