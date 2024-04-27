package face

import "github.com/funtimecoding/go-library/pkg/web/web_client"

type WebClient interface {
	IsWebClient()

	Post(
		locator string,
		body any,
	) (*web_client.Response, error)
}
