package mock_snippet

func New() *Client {
	return &Client{files: make(map[string][]byte)}
}
