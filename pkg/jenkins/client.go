package jenkins

import (
	"context"
	"github.com/bndr/gojenkins"
	"github.com/funtimecoding/go-library/pkg/jenkins/basic"
)

type Client struct {
	basic   *basic.Client
	context context.Context
	client  *gojenkins.Jenkins
}
