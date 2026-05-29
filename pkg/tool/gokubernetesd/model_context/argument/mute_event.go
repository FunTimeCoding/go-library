package argument

type MuteEvent struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Cluster string `json:"cluster"`
}
