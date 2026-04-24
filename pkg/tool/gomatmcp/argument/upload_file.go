package argument

type UploadFile struct {
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Path        string `json:"path"`
	Message     string `json:"message"`
}
