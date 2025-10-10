package locator

import "net/url"

type Locator struct {
	values   url.Values
	scheme   string
	host     string
	port     int
	basePath string
	path     string
	fragment string
}
