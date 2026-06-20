package runner

type SyncRequest struct {
	Response chan *SyncResult
}
