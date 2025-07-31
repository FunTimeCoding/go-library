package repository

import (
	"github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"strings"
)

func New(
	path string,
	status string,
) *Repository {
	return &Repository{
		MonitorIdentifier: constant.GoGitLab.StringIdentifier(path),
		Path:              path,
		IsClean:           len(strings.TrimSpace(status)) == 0,
		Status:            status,
	}
}
