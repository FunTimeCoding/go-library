package option

import "time"

type Exporter struct {
	Owner    string
	Interval time.Duration
	Verbose  bool
}
