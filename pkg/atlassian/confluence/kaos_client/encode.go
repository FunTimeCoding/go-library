package kaos_client

import (
	"github.com/funtimecoding/go-library/pkg/strings/base64"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/web/constant"
)

func (a *Credential) Encode() string {
	return key_value.Space(
		constant.BasicPrefix,
		base64.Encode(key_value.Colon(a.User, a.Token)),
	)
}
