package profile

import (
	"github.com/funtimecoding/go-library/pkg/brave/helper"
	"github.com/funtimecoding/go-library/pkg/strings/base64"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func Parse(
	base string,
	profile string,
	s string,
) *Profile {
	_, mail, _, _ := helper.ParseAccounts([]byte(base64.Decode(s)))
	path := join.Absolute(base, profile)

	return &Profile{Profile: profile, Path: path, Email: mail}
}
