package runner

import "fmt"

func (r *Runner) Sync() (*SyncResult, error) {
	request := SyncRequest{Response: make(chan *SyncResult, 1)}

	select {
	case r.sync <- request:
		return <-request.Response, nil
	default:
		return nil, fmt.Errorf("sync already queued")
	}
}
