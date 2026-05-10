package response

type KeysReturn struct {
	Local           []string `json:"local"`
	Minions         []string `json:"minions"`
	MinionsPre      []string `json:"minions_pre"`
	MinionsRejected []string `json:"minions_rejected"`
	MinionsDenied   []string `json:"minions_denied"`
}
