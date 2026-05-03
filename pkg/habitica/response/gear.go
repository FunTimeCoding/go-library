package response

type Gear struct {
	Equipped map[string]string `json:"equipped"`
	Owned    map[string]bool   `json:"owned"`
}
