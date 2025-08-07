package post

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/text"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
)

func (p *Post) Format() string {
	return fmt.Sprintf(
		"%s %s: %s",
		p.Create.Format(timeLibrary.HourMinute),
		p.User.Username,
		text.OptimizeWhitespace(p.Message, constant.CompactWhitespace),
	)
}
