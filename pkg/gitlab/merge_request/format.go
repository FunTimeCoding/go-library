package merge_request

import (
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
)

func (m *Request) Format(s *format.Settings) string {
	return status.New(s).Integer(
		m.Project,
		m.Identifier,
	).String(m.Title, m.State).Raw(m.Raw).Format()
}
