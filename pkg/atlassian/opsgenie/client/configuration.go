package client

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/sirupsen/logrus"
)

func Configuration(key string) *client.Config {
	return &client.Config{ApiKey: key, LogLevel: logrus.ErrorLevel}
}
