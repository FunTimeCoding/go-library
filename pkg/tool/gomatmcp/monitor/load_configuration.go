package monitor

import (
	library "github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"os"
)

func LoadConfiguration() *Configuration {
	if !environment.Exists("MONITORING_ENABLED") {
		return &Configuration{Enabled: false}
	}

	return &Configuration{
		Enabled:  true,
		Schedule: environment.Required("MONITORING_SCHEDULE"),
		Channels: parseList(
			environment.Required("MONITORING_CHANNELS"),
		),
		Topics: parseList(
			environment.Required("MONITORING_TOPICS"),
		),
		MessageLimit: library.ToIntegerStrict(
			environment.Required("MONITORING_MESSAGE_LIMIT"),
		),
		NotificationChannel: os.Getenv("MONITORING_NOTIFICATION_CHANNEL"),
		Username:            os.Getenv("MONITORING_USERNAME"),
	}
}
