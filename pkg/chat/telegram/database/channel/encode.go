package channel

import "github.com/funtimecoding/go-library/pkg/notation"

func (c *Channel) Encode() []byte {
	return notation.Marshall(c)
}
