package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/system"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"github.com/prometheus/alertmanager/api/v2/client/silence"
	"github.com/prometheus/alertmanager/api/v2/models"
	"log"
	"time"
)

func (c *Client) CreateSilence(
	alert string,
	comment string,
	d time.Duration,
) *silence.PostSilencesOKBody {
	r := c.Rules().Find(alert)

	if r == nil {
		log.Panicf("alert not found: %s", alert)
	}

	if comment == "" {
		comment = constant.NoComment
	}

	t := time.Now()

	if d == 0 {
		d = 10 * time.Minute
	}

	p := silence.NewPostSilencesParams()
	p.Silence = &models.PostableSilence{
		Silence: models.Silence{
			Comment:   &comment,
			CreatedBy: ptr.To(system.User().Username),
			Matchers: []*models.Matcher{
				{
					Name:    ptr.To(constant.NameField),
					Value:   &alert,
					IsRegex: ptr.To(false),
				},
			},
			StartsAt: ptr.To(timeLibrary.Scan(t)),
			EndsAt:   ptr.To(timeLibrary.Scan(t.Add(d))),
		},
	}

	result, e := c.client.Silence.PostSilences(p)
	errors.PanicOnError(e)

	return result.Payload
}
