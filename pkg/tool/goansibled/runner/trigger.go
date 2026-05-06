package runner

import "fmt"

func (r *Runner) Trigger(playbook []string) error {
	request := TriggerRequest{Playbook: playbook}

	select {
	case r.trigger <- request:
		return nil
	default:
		return fmt.Errorf("run already queued")
	}
}
