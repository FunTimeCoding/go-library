package sublime

import (
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"net/http"
)

type Client struct {
	base   *locator.Locator
	client *http.Client
}
