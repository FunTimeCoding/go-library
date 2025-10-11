package locator

import "net/url"

type Locator struct {
	scheme        string
	user          string
	password      string
	host          string
	port          int
	basePath      string
	path          string
	trail         bool
	value         url.Values
	fragment      string
	fragmentValue url.Values
}
