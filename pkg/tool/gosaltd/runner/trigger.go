package runner

import "fmt"

func (r *Runner) Trigger() error {
	select {
	case r.trigger <- TriggerRequest{}:
		return nil
	default:
		return fmt.Errorf("run already queued")
	}
}
