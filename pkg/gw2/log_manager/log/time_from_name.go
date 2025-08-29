package log

import (
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
	library "github.com/funtimecoding/go-library/pkg/time"
	"strings"
	"time"
)

func timeFromName(fileName string) time.Time {
	parts := strings.Split(fileName, "\\")
	date, _ := key_value.Dot(parts[len(parts)-1])

	return library.Parse("20060102-150405", date)
}
