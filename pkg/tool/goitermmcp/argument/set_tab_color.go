package argument

import "github.com/funtimecoding/go-library/pkg/generative/mark/request"

type SetTabColor struct {
	SessionIdentifier string          `json:"session_id"`
	Red               request.Integer `json:"red"`
	Green             request.Integer `json:"green"`
	Blue              request.Integer `json:"blue"`
}
