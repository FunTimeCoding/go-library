package response

import "time"

type Version struct {
	Number           int       `json:"number"`
	Message          string    `json:"message"`
	MinorEdit        bool      `json:"minorEdit"`
	AuthorIdentifier string    `json:"authorId"`
	CreatedAt        time.Time `json:"createdAt"`
}
