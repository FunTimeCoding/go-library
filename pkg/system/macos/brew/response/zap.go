package response

type Zap struct {
	Trash           any      `json:"trash"`
	RemoveDirectory any      `json:"rmdir,omitempty"`
	LaunchControl   []string `json:"launchctl,omitempty"`
}
