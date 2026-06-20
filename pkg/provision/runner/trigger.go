package runner

import "fmt"

func (r *Runner) Trigger(request TriggerRequest) error {
	select {
	case r.trigger <- request:
		return nil
	default:
		return fmt.Errorf("run already queued")
	}
}
