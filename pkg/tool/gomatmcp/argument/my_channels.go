package argument

type MyChannels struct {
	Limit      int    `json:"limit"`
	TypeFilter string `json:"type"`
	Since      string `json:"since"`
}
