package mock_connection

func (c Connection) Write(_ []byte) (n int, err error) {
	return 0, nil
}
