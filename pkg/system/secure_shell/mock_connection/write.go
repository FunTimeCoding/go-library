package mock_connection

func (c Connection) Write(_ []byte) (n int, e error) {
	return 0, nil
}
