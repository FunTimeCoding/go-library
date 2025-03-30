package response

import "time"

type Space struct {
	AuthorId           string    `json:"authorId"`
	HomepageId         string    `json:"homepageId"`
	Icon               string    `json:"icon"`
	Description        string    `json:"description"`
	Status             string    `json:"status"`
	CreatedAt          time.Time `json:"createdAt"`
	Name               string    `json:"name"`
	Key                string    `json:"key"`
	Id                 string    `json:"id"`
	Type               string    `json:"type"`
	Links              Links     `json:"_links"`
	CurrentActiveAlias string    `json:"currentActiveAlias"`
}
