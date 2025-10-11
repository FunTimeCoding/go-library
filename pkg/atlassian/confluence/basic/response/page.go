package response

import "time"

type Page struct {
	SpaceId                string    `json:"spaceId"`
	ParentId               string    `json:"parentId"`
	OwnerId                string    `json:"ownerId"`
	LastOwnerId            string    `json:"lastOwnerId"`
	AuthorId               string    `json:"authorId"`
	ParentType             string    `json:"parentType"`
	Version                Version   `json:"version"`
	Position               int       `json:"position"`
	Status                 string    `json:"status"`
	Body                   Body      `json:"body"`
	CreatedAt              time.Time `json:"createdAt"`
	Title                  string    `json:"title"`
	Id                     string    `json:"id"`
	Links                  Links     `json:"_links"`
	SourceTemplateEntityId string    `json:"sourceTemplateEntityId,omitempty"`
}
