package minute

import (
	"github.com/funtimecoding/go-library/pkg/time"
	"github.com/funtimecoding/go-library/pkg/time/second"
)

func Readable(minutes int) string {
	return second.Readable(minutes * time.MinuteInSeconds)
}
