package item

import "time"

type Item struct {
	Identifier string
	Type       string
	Detail     string
	Link       string
	Create     *time.Time
}
