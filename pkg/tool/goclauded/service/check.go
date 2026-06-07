package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/service/check_result"

func (s *Service) Check(sessionIdentifier string) (*check_result.Result, error) {
	r, e := s.store.EnsureSession(sessionIdentifier)

	if e != nil {
		return nil, e
	}

	roster, e := s.store.ConsumeRoster(r.Callsign)

	if e != nil {
		return nil, e
	}

	changes, e := s.store.HasChanges(r.Callsign, r.LastSeen)

	if e != nil {
		return nil, e
	}

	changed := roster || changes
	result := check_result.New()
	result.Callsign = r.Callsign
	pulses, e := s.store.PendingPulses(sessionIdentifier)

	if e != nil {
		return nil, e
	}

	result.Pulses = pulses

	if len(result.Pulses) > 0 {
		result.Changed = true
	}

	if changed {
		result.Changed = true
		sessions, e := s.store.ListSessions()

		if e != nil {
			return nil, e
		}

		result.Sessions = sessions
		messages, e := s.store.PendingMessages(r.Callsign)

		if e != nil {
			return nil, e
		}

		result.Messages = messages
		completions, e := s.store.RecentCompletions()

		if e != nil {
			return nil, e
		}

		result.Completions = completions
	}

	if !r.LastSeen.IsZero() {
		since := r.LastSeen.UTC().Format("2006-01-02T15:04:05Z")
		activity := s.MemoryActivity(since)

		if len(activity) > 0 {
			result.MemoryActivity = activity
			result.Changed = true
		}
	}

	timeout, e := s.store.ConsumeTimeout(r.Callsign)

	if e != nil {
		return nil, e
	}

	if timeout != "" {
		result.TimeoutMessage = timeout
		result.Changed = true
	}

	reannounce, e := s.store.ConsumeReannounce(r.Callsign)

	if e != nil {
		return nil, e
	}

	if reannounce {
		result.Reannounce = true
		result.Changed = true
	}

	s.notify()

	return result, nil
}
