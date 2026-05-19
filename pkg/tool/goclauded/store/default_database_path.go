package store

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"path/filepath"
)

func DefaultDatabasePath() string {
	n := constant.Identity.Name()
	d := filepath.Join(system.Home(), ".local", "share", n)
	system.MakeDirectory(d)

	return filepath.Join(d, fmt.Sprintf("%s.sqlite", n))
}
