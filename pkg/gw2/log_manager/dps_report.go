package log_manager

type DpsReport struct {
	Locator         any `json:"Url"`
	ProcessingError any `json:"ProcessingError"`
	UploadTime      any `json:"UploadTime"`
}
