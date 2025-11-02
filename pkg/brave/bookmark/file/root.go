package file

type Root struct {
	Bar    *Node `json:"bookmark_bar"`
	Other  *Node `json:"other"`
	Synced *Node `json:"synced"`
}
