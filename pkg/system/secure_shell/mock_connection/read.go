package mock_connection

func (c Connection) Read(_ []byte) (n int, e error) {
	return 0, nil
}
