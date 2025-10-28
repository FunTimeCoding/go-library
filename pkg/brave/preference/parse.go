package preference

import (
	"github.com/funtimecoding/go-library/pkg/brave/constant"
	"github.com/funtimecoding/go-library/pkg/brave/helper"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func Parse(profile string) *Preference {
	var result Preference
	p := join.Absolute(helper.SettingsPath(), profile)
	b := system.ReadBytes(p, constant.PreferencesFile)
	notation.DecodeBytesStrict(b, &result, false)

	return &result
}
