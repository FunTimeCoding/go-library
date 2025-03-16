package silence

import (
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

type Silence struct {
	MonitorIdentifier string
	Identifier        string
	State             string
	Match             string
	Author            string
	Comment           string

	Start *time.Time
	End   *time.Time

	Rule string
	Link string

	Raw *models.GettableSilence
}
