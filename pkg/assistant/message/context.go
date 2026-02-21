package message

type Context struct {
	Identifier string `json:"id"`
	Parent     string `json:"parent_id,omitempty"`
	User       string `json:"user_id,omitempty"`
}
