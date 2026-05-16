package response

type Project struct {
	Self           string        `json:"self"`
	Identifier     string        `json:"id"`
	Key            string        `json:"key"`
	Name           string        `json:"name"`
	ProjectTypeKey string        `json:"projectTypeKey"`
	Simplified     bool          `json:"simplified"`
	AvatarLocator  AvatarLocator `json:"avatarUrls"`
}
