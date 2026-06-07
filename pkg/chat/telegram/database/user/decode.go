package user

import "github.com/funtimecoding/go-library/pkg/notation"

func (u *User) Decode(b []byte) {
	notation.MustDecodeBytes(b, u, false)
}
