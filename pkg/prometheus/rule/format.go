package rule

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"slices"
	"time"
)

func (r *Rule) Format(s *format.Settings) string {
	t := status.New(s).String(
		r.formatName(s),
		r.Group,
		r.formatType(),
	).Raw(r)

	if r.Summary != "" {
		t.Line("  Summary: %s", r.Summary)
	}

	if r.Description != "" {
		t.Line("  Description: %s", r.Description)
	}

	if r.Duration > 0 {
		t.Line("  Duration: %d", r.Duration)
	}

	if s.ShowRaw && r.RawAlert != nil {
		t.Line("  RawAlert: %+v", r.RawAlert)
	}

	if !slices.Contains(Healths, r.Health) {
		output := fmt.Sprintf("Unexpected health: %s", r.Health)

		if s.UseColor {
			output = console.Red(output)
		}

		t.Line("%s", output)
	}

	if !slices.Contains(States, r.State) {
		output := fmt.Sprintf("Unexpected state: %s", r.State)

		if s.UseColor {
			output = console.Red(output)
		}

		t.Line("%s", output)
	}

	if r.RawAlert != nil {
		if age := time.Since(
			r.RawAlert.LastEvaluation,
		).Round(time.Second); age > 30*time.Second {
			output := fmt.Sprintf("  Alert last evaluation: %s", age)

			if s.UseColor {
				output = console.Red(output)
			}

			t.Line("%s", output)
		}

		if r.RawAlert.EvaluationTime > 0.1 {
			output := fmt.Sprintf(
				"  EvaluationTime: %.1f",
				r.RawAlert.EvaluationTime,
			)

			if s.UseColor {
				output = console.Yellow(output)
			}

			t.Line("%s", output)
		}

		if false {
			if false {
				r.RawAlert.LastEvaluation = time.Time{}
			}

			r.RawAlert.Name = ""
			r.RawAlert.State = ""
			r.RawAlert.Health = ""
			r.RawAlert.Duration = 0
			r.RawAlert.Query = ""
			delete(r.RawAlert.Labels, SeverityKey)
			delete(r.RawAlert.Annotations, SummaryKey)
			delete(r.RawAlert.Annotations, DescriptionKey)
			delete(r.RawAlert.Annotations, RunbookKey)
			delete(r.RawAlert.Annotations, DurationKey)
			t.Line("  RawAlert: %+v", r.RawAlert)
		}
	}

	if s.ShowRaw && r.RawRecord != nil {
		t.Line("  RawRecord: %+v", r.RawRecord)
	}

	return t.Format()
}
