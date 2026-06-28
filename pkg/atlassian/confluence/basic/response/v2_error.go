package response

type V2Error struct {
	Errors []V2ErrorEntry `json:"errors"`
}
