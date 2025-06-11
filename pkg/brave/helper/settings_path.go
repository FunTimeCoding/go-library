package helper

import (
	"github.com/funtimecoding/go-library/pkg/brave/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"path/filepath"
)

func SettingsPath() string {
	return filepath.Join(system.Home(), constant.BraveSettings)
}
