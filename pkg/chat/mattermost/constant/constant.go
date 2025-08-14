package constant

import "github.com/funtimecoding/go-library/pkg/console/status/option"

const (
	HostEnvironment    = "MATTERMOST_HOST"
	TokenEnvironment   = "MATTERMOST_TOKEN"
	TeamEnvironment    = "MATTERMOST_TEAM"
	ChannelEnvironment = "MATTERMOST_CHANNEL"

	PerPage int = 1000

	EmptyEntityTag = ""
)

// Emoji
const (
	CheckMark    = "white_check_mark" // done
	Construction = "construction"     // in progress
	Repeat       = "repeat"           // forwarded
)

var Format = option.ExtendedColor.Copy()
