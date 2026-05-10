package response

type LocalReturn struct {
	Jid     string `json:"jid"`
	Ret     any    `json:"ret"`
	Retcode int    `json:"retcode"`
}
