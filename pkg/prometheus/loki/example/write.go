package example

import (
	hook "github.com/akkuman/logrus-loki-hook"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"github.com/sirupsen/logrus"
)

func Write() {
	host := environment.Required(constant.HostEnvironment)
	user := environment.Required(constant.UserEnvironment)
	password := environment.Required(constant.PasswordEnvironment)
	l := logrus.New()
	h, e := hook.NewHook(
		&hook.Config{
			URL: locator.New(
				host,
			).UserPassword(user, password).Path("/api/prom/push").String(),
			LevelName: "severity",
			Labels:    map[string]string{"application": "test"},
		},
	)

	if e != nil {
		l.Error(e)
	} else {
		l.AddHook(h)
	}

	l.Info("test message")
}
