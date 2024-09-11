package gitlab

func NewGitLabCom(
	token string,
	projects []int,
) *Client {
	return New("", token, projects)
}
