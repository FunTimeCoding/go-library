package response

import "time"

type User struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	IsActive   bool       `json:"isActive"`
	DateJoined *time.Time `json:"dateJoined"`
}
