package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request"

type SaveView struct {
	ViewIdentifier request.Integer `json:"view_id"`
	FilePath       string          `json:"file_path"`
}
