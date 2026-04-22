package convert

import "github.com/funtimecoding/go-library/pkg/tool/goatld/server"

type UpdateResult struct {
	Issue   *server.JiraIssue `json:"issue"`
	Changes []FieldChange     `json:"changes,omitempty"`
}
