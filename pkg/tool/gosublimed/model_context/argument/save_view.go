package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request/integer"

type SaveView struct {
	ViewIdentifier integer.Integer `json:"view_id"`
	FilePath       string          `json:"file_path"`
}
