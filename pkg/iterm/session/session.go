package session

type Session struct {
	Identifier       string `json:"session_id"`
	WindowIdentifier string `json:"window_id"`
	WindowTitle      string `json:"window_title"`
	TabIdentifier    string `json:"tab_id"`
	TabTitle         string `json:"tab_title"`
	JobName          string `json:"job_name"`
	CommandLine      string `json:"command_line"`
	Path             string `json:"path"`
	Tty              string `json:"tty"`
	Columns          int    `json:"columns"`
	Rows             int    `json:"rows"`
}
