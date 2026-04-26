package issue

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	library "github.com/funtimecoding/go-library/pkg/time"
	"time"
)

func ConvertTime(s string) time.Time {
	return library.Parse(constant.TimeFormat, s)
}
