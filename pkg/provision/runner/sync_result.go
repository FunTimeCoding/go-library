package runner

type SyncResult struct {
	Changed bool   `json:"changed"`
	Diff    string `json:"diff,omitempty"`
	Error   error  `json:"-"`
}
