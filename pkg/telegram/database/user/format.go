package user

import "fmt"

func (u *User) Format() string {
	return fmt.Sprintf("%d %s\n", u.Identifier, u.Name)
}
