package node_status

type Swap struct {
	Free  int64 `json:"free"`
	Total int64 `json:"total"`
	Used  int   `json:"used"`
}
