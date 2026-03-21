package mock_connection

import "time"

func (c Connection) SetWriteDeadline(_ time.Time) error {
	return nil
}
