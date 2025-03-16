package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/system"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"github.com/prometheus/alertmanager/api/v2/client/silence"
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

func (c *Client) PostSilence(
	identifier string,
	alert string,
	comment string,
	start time.Time,
	end time.Time,
) string {
	if comment == "" {
		comment = constant.NoComment
	}

	p := silence.NewPostSilencesParams()
	p.Silence = &models.PostableSilence{
		ID: identifier,
		Silence: models.Silence{
			Comment:   &comment,
			CreatedBy: ptr.To(system.User().Username),
			Matchers: []*models.Matcher{
				{
					Name:    ptr.To(constant.AlertnameField),
					Value:   &alert,
					IsRegex: ptr.To(false),
				},
			},
			StartsAt: ptr.To(timeLibrary.Scan(start)),
			EndsAt:   ptr.To(timeLibrary.Scan(end)),
		},
	}

	result, e := c.client.Silence.PostSilences(p)
	errors.PanicOnError(e)

	return result.Payload.SilenceID
}
