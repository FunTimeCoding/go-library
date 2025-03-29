package user

import "github.com/opsgenie/opsgenie-go-sdk-v2/user"

func NewSlice(v []user.User) []*User {
	var result []*User

	for _, element := range v {
		result = append(result, New(&element))
	}

	return result
}
