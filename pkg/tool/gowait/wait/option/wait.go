package option

import "time"

type Wait struct {
	Locator  string
	Contains string
	File     string
	Process  string
	Verbose  bool
	Timeout  time.Duration
}
