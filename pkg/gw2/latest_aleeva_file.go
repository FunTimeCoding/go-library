package gw2

import (
	"fmt"
	"strings"
	"time"
)

func LatestAleevaFile(files []string) string {
	var latest time.Time

	for _, file := range files {
		date := strings.TrimPrefix(file, AleevaPrefix)
		date = strings.TrimSuffix(date, NotationSuffix)
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
		AleevaPrefix,
		latest.Format("2006-01-02"),
		NotationSuffix,
	)
}
