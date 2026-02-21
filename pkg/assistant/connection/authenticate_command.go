package connection

type authenticateCommand struct {
	Type  string `json:"type"`
	Token string `json:"access_token"`
}
