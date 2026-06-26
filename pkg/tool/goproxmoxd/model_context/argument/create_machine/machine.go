package create_machine

type Machine struct {
	Node            string `json:"node"`
	Name            string `json:"name"`
	Identifier      int    `json:"identifier"`
	Cores           int    `json:"cores"`
	Memory          int    `json:"memory"`
	DiskSize        int    `json:"disk_size"`
	DiskStorage     string `json:"disk_storage"`
	DiskImport      string `json:"disk_import"`
	CDROM           string `json:"cdrom"`
	Bridge          string `json:"bridge"`
	OSType          string `json:"os_type"`
	Agent           *bool  `json:"agent"`
	Start           bool   `json:"start"`
	Tags            string `json:"tags"`
	CIUser          string `json:"ci_user"`
	CIPassword      string `json:"ci_password"`
	SSHKeys         string `json:"ssh_keys"`
	IPConfiguration string `json:"ip_config"`
	Extras          string `json:"extras"`
}
