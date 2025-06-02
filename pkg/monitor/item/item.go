package item

import "time"

type Item struct {
	Identifier string
	Level      string
	Detail     string
	Link       string
	Create     *time.Time
}
