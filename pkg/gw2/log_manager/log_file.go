package log_manager

type LogFile struct {
	Version        int            `json:"Version"`
	LogsByFilename map[string]any `json:"LogsByFilename"`
}
