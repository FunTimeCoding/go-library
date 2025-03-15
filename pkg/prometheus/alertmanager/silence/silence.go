package silence

import (
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

type Silence struct {
	Identifier string
	State      string
	Match      string
	Author     string
	Comment    string

	Start *time.Time
	End   *time.Time

	Raw *models.GettableSilence
}
