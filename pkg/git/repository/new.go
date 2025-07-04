package repository

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"strings"
)

func New(
	path string,
	status string,
) *Repository {
	return &Repository{
		MonitorIdentifier: fmt.Sprintf(
			"%s-%s",
			constant.GitPrefix,
			path,
		),
		Path:    path,
		IsClean: len(strings.TrimSpace(status)) == 0,
		Status:  status,
	}
}
