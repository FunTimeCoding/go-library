package jenkins

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/jenkins/basic_client"
	"github.com/funtimecoding/go-library/pkg/web/constant"
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
	locator := fmt.Sprintf(
		"%s://%s:%d/",
		constant.InsecureScheme,
		host,
		port,
	)
	client := gojenkins.CreateJenkins(
		nil,
		locator,
		user,
		password,
	)
	c := context.Background()
	_, e := client.Init(c)
	errors.PanicOnError(e)

	return &Client{
		basic:   basic_client.New(locator, user, password),
		context: c,
		client:  client,
	}
}
