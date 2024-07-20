package log

import (
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"strings"
	"time"
)

func timeFromName(fileName string) time.Time {
	parts := strings.Split(fileName, "\\")
	date, _ := key_value.Dot(parts[len(parts)-1])

	return timeLibrary.Parse("20060102-150405", date)
}
