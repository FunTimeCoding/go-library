package client

func Stub() *Client {
	return &Client{clients: make(map[string]*Client)}
}
