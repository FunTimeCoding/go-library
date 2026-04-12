package option

import "time"

type Loki struct {
	Namespace  string
	Since      time.Duration
	Route      string
	Message    string
	BodyOnly   bool
	Copyable   bool
	Namespaces []string
	Exclude    []string
	Limit      int
}
