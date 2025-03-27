package kaos_client

import (
	"encoding/base64"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (a *Credential) Encode() string {
	return key_value.Space(
		web.BasicPrefix,
		base64.StdEncoding.EncodeToString(
			[]byte(key_value.Colon(a.User, a.Token)),
		),
	)
}
