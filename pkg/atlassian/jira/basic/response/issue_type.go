package response

type IssueType struct {
	Self             string `json:"self"`
	Identifier       string `json:"id"`
	Description      string `json:"description"`
	IconLocator      string `json:"iconUrl"`
	Name             string `json:"name"`
	Subtask          bool   `json:"subtask"`
	AvatarIdentifier int    `json:"avatarId"`
	HierarchyLevel   int    `json:"hierarchyLevel"`
}
