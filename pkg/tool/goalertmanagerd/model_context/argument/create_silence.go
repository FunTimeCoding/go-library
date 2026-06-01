package argument

type CreateSilence struct {
	Alert    string `json:"alert"`
	Comment  string `json:"comment"`
	Duration string `json:"duration"`
}
