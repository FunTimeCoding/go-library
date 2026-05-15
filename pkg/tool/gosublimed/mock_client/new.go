package mock_client

func New() *Client {
	return &Client{nextID: 1}
}
