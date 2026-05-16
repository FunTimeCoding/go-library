package gear

type Gear struct {
	Equipped map[string]string `json:"equipped"`
	Owned    map[string]bool   `json:"owned"`
}
