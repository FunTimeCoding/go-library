package channel

import "github.com/funtimecoding/go-library/pkg/notation"

func (c *Channel) Decode(b []byte) {
	notation.MustDecodeBytes(b, c, false)
}
