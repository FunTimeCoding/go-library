package gitlab

func NewGitLabCom(
	token string,
	o ...Option,
) *Client {
	return New("", token, o...)
}
