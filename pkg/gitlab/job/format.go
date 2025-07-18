package job

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (j *Job) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer(j.Identifier)
	}

	s.String(j.formatName(f))

	if f.HasTag(tag.Project) {
		s.String(j.formatProject())
	}

	s.String(
		j.formatUser(),
		j.formatDate(f),
		j.formatConcern(f),
	)

	if f.HasTag(tag.Link) {
		s.TagLine(tag.Link, "  %s", j.Link)
		s.Line("  Project: %s", j.ProjectLink())
	}

	return s.RawList(j.Raw).Format()
}
