package mock_connection

import "time"

func (c Connection) SetReadDeadline(_ time.Time) error {
	return nil
}
