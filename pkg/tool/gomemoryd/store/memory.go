package store

type Memory struct {
	Identifier       int64    `json:"identifier"`
	Name             string   `json:"name"`
	Content          string   `json:"content"`
	Description      string   `json:"description"`
	Type             string   `json:"type"`
	CreatedAt        string   `json:"created_at"`
	UpdatedAt        string   `json:"updated_at"`
	IsActive         bool     `json:"is_active"`
	Tags             []string `json:"tags,omitempty"`
	ParentIdentifier *int64   `json:"parent_identifier,omitempty"`
}
