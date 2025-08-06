package user

import "github.com/funtimecoding/go-library/pkg/notation"

func (u *User) Encode() []byte {
	return notation.Marshall(u)
}
