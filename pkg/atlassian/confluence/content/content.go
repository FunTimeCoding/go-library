package content

import (
	kaos "github.com/essentialkaos/go-confluence/v6"
	virtomize "github.com/virtomize/confluence-go-api"
)

type Content struct {
	Identifier string
	Link       string
	TinyLink   string
	Space      string
	Title      string
	Labels     []string

	Raw          *kaos.Content
	RawVirtomize *virtomize.Results
}
