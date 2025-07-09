package gw2

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gw2/constant"
	"strings"
	"time"
)

func LatestAleevaFile(files []string) string {
	var latest time.Time

	for _, file := range files {
		date := strings.TrimPrefix(file, constant.AleevaPrefix)
		date = strings.TrimSuffix(date, constant.NotationSuffix)
		t, err := time.Parse("2006-01-02", date)

		if err != nil {
			continue
		}

		if t.After(latest) {
			latest = t
		}
	}

	return fmt.Sprintf(
		"%s%s%s",
		constant.AleevaPrefix,
		latest.Format("2006-01-02"),
		constant.NotationSuffix,
	)
}
