package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"goraidd",
	"RAID monitoring service",
	"goraidd [flags]",
)

const (
	LogsTitle    = "Logs"
	LogsPath     = "/"
	ReportsTitle = "Reports"
	ReportsPath  = "/reports"
	RaidsTitle   = "Raids"
	RaidsPath    = "/raids"
	PlayersTitle = "Players"
	PlayersPath  = "/players"
)
