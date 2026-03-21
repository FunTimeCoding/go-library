package mock_connection

import "time"

func (c Connection) SetDeadline(_ time.Time) error {
	return nil
}
