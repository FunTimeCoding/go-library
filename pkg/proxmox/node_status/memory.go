package node_status

type Memory struct {
	Available int64 `json:"available"`
	Free      int64 `json:"free"`
	Total     int64 `json:"total"`
	Used      int64 `json:"used"`
}
