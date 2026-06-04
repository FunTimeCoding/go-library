package service

func (s *Service) summaryMetadata(
	sessionIdentifier string,
	name string,
) map[string]string {
	result := map[string]string{
		"session_name": name,
		"session_id":   sessionIdentifier,
	}
	session := s.store.GetSession(sessionIdentifier)

	if session != nil && !session.StartedAt.IsZero() {
		result["date"] = session.StartedAt.Format("2006-01-02")
	}

	return result
}
