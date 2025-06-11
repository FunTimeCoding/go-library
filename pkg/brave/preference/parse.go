package preference

import (
	"github.com/funtimecoding/go-library/pkg/brave/constant"
	"github.com/funtimecoding/go-library/pkg/brave/helper"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system"
	"path/filepath"
)

func Parse(profile string) *Preference {
	var result Preference
	notation.DecodeBytes(
		system.ReadBytes(
			filepath.Join(
				helper.SettingsPath(),
				profile,
				constant.PreferencesFile,
			),
		),
		&result,
	)

	return &result
}
