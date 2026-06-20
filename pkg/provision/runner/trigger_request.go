package runner

type TriggerRequest struct {
	Parameters map[string]any
	Update     bool
	Response   chan *TriggerResult
}
