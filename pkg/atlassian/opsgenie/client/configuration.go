package client

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/sirupsen/logrus"
	"time"
)

func Configuration(key string) *client.Config {
	return &client.Config{
		ApiKey:         key,
		LogLevel:       logrus.ErrorLevel,
		RetryCount:     2,
		RequestTimeout: 10 * time.Second,
	}
}
