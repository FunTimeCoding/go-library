package response

import "time"

type Page struct {
	SpaceIdentifier                string    `json:"spaceId"`
	ParentIdentifier               string    `json:"parentId"`
	OwnerIdentifier                string    `json:"ownerId"`
	LastOwnerIdentifier            string    `json:"lastOwnerId"`
	AuthorIdentifier               string    `json:"authorId"`
	ParentType                     string    `json:"parentType"`
	Version                        Version   `json:"version"`
	Position                       int       `json:"position"`
	Status                         string    `json:"status"`
	Body                           Body      `json:"body"`
	CreatedAt                      time.Time `json:"createdAt"`
	Title                          string    `json:"title"`
	Identifier                     string    `json:"id"`
	Links                          Links     `json:"_links"`
	SourceTemplateEntityIdentifier string    `json:"sourceTemplateEntityId,omitempty"`
}
