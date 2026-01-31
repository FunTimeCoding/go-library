package constant

import "github.com/funtimecoding/go-library/pkg/console/status/option"

const (
	HostEnvironment    = "MATTERMOST_HOST"
	TokenEnvironment   = "MATTERMOST_TOKEN"
	TeamEnvironment    = "MATTERMOST_TEAM"
	ChannelEnvironment = "MATTERMOST_CHANNEL"

	PerPage int = 1000

	EmptyEntityTag = ""

	PostField = "post"
)

// Emoji
const (
	CheckMark    = "white_check_mark"       // done
	Construction = "construction"           // in progress
	Hourglass    = "hourglass_flowing_sand" // waiting, pending resolve
	Repeat       = "repeat"                 // forwarded
	Thread       = "thread"                 // belongs to above thread
)

var Format = option.ExtendedColor.Copy()
