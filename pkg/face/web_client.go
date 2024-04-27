package face

import "github.com/funtimecoding/go-library/pkg/web/web_client/web_response"

type WebClient interface {
	IsWebClient()

	Post(
		locator string,
		body any,
	) (*web_response.Response, error)
}
