package argument

type DownloadFile struct {
	FileIdentifier string `json:"file_id"`
	Path           string `json:"path"`
}
