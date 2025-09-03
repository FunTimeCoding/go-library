package jenkins

import (
	"context"
	"github.com/bndr/gojenkins"
	"github.com/funtimecoding/go-library/pkg/jenkins/basic_client"
)

type Client struct {
	basic   *basic_client.Client
	context context.Context
	client  *gojenkins.Jenkins
}
