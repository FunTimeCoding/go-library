package response

import "encoding/json"

type LoginReturn struct {
	Token  string          `json:"token"`
	Expire float64         `json:"expire"`
	Start  float64         `json:"start"`
	User   string          `json:"user"`
	EAuth  string          `json:"eauth"`
	Perms  json.RawMessage `json:"perms"`
}
