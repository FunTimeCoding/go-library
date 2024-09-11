package gitlab

func NewGitLabCom(token string) *Client {
	return New("", token)
}
