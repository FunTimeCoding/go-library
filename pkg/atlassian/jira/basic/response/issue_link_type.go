package response

type IssueLinkType struct {
	Identifier string `json:"id"`
	Name       string `json:"name"`
	Inward     string `json:"inward"`
	Outward    string `json:"outward"`
	Self       string `json:"self"`
}
