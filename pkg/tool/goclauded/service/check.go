package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/service/check_result"

func (s *Service) Check(sessionIdentifier string) *check_result.Result {
	r := s.store.EnsureSession(sessionIdentifier)
	changed := s.store.ConsumeRoster(r.Callsign) ||
		s.store.HasChanges(r.Callsign, r.LastSeen)
	result := check_result.New()
	result.Callsign = r.Callsign
	result.Pulses = s.store.PendingPulses(sessionIdentifier)

	if len(result.Pulses) > 0 {
		result.Changed = true
	}

	if changed {
		result.Changed = true
		result.Sessions = s.store.ListSessions()
		result.Messages = s.store.PendingMessages(r.Callsign)
		result.Completions = s.store.RecentCompletions()
	}

	if !r.LastSeen.IsZero() {
		since := r.LastSeen.UTC().Format("2006-01-02T15:04:05Z")
		activity := s.MemoryActivity(since)

		if len(activity) > 0 {
			result.MemoryActivity = activity
			result.Changed = true
		}
	}

	if timeout := s.store.ConsumeTimeout(r.Callsign); timeout != "" {
		result.TimeoutMessage = timeout
		result.Changed = true
	}

	if s.store.ConsumeReannounce(r.Callsign) {
		result.Reannounce = true
		result.Changed = true
	}

	s.notify()

	return result
}
