package node_status

type RootFilesystem struct {
	Available int64 `json:"avail"`
	Free      int64 `json:"free"`
	Total     int64 `json:"total"`
	Used      int64 `json:"used"`
}
