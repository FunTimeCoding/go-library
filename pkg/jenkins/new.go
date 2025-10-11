package jenkins

import (
	"context"
	"github.com/bndr/gojenkins"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/jenkins/basic"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func New(
	host string,
	user string,
	password string,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(user, "user")
	errors.FatalOnEmpty(password, "password")
	port := 8080
	client := gojenkins.CreateJenkins(
		nil,
		locator.New(host).Port(port).Insecure().String(),
		user,
		password,
	)
	c := context.Background()
	_, e := client.Init(c)
	errors.PanicOnError(e)

	return &Client{
		basic:   basic.New(host, port, user, password),
		context: c,
		client:  client,
	}
}
