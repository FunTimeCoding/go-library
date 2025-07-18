package gitlab

func NewGitLabCom(
	token string,
	o ...OptionFunc,
) *Client {
	return New("", token, o...)
}
