package response

type Uninstall struct {
	Quit            string          `json:"quit,omitempty"`
	LoginItem       string          `json:"login_item,omitempty"`
	PackageUtility  any             `json:"pkgutil,omitempty"`
	LaunchControl   any             `json:"launchctl,omitempty"`
	Script          UninstallScript `json:"script,omitempty"`
	Delete          any             `json:"delete,omitempty"`
	RemoveDirectory string          `json:"rmdir,omitempty"`
	Signal          []string        `json:"signal,omitempty"`
	Trash           string          `json:"trash,omitempty"`
}
