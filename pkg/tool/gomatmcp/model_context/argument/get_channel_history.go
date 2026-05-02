package argument

type GetChannelHistory struct {
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Limit       int    `json:"limit"`
	Since       string `json:"since"`
}
