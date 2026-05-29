package argument

type ExecInPod struct {
	Name      string   `json:"name"`
	Namespace string   `json:"namespace"`
	Command   []string `json:"command"`
	Container string   `json:"container"`
	Timeout   int      `json:"timeout"`
}
