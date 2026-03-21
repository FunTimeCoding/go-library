package mock_connection

func (c Connection) Read(_ []byte) (n int, err error) {
	return 0, nil
}
