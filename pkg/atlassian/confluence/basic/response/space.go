package response

import "time"

type Space struct {
	AuthorIdentifier   string    `json:"authorId"`
	HomepageIdentifier string    `json:"homepageId"`
	Icon               string    `json:"icon"`
	Description        string    `json:"description"`
	Status             string    `json:"status"`
	CreatedAt          time.Time `json:"createdAt"`
	Name               string    `json:"name"`
	Key                string    `json:"key"`
	Identifier         string    `json:"id"`
	Type               string    `json:"type"`
	Links              Links     `json:"_links"`
	CurrentActiveAlias string    `json:"currentActiveAlias"`
}
