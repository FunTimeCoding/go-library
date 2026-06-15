package service

func (s *Service) resolveAliases(sessionIDs map[string]bool) map[string]string {
	result := map[string]string{}

	for identifier := range sessionIDs {
		session, e := s.store.GetSession(identifier)

		if e != nil || session == nil {
			continue
		}

		if session.Alias != nil && *session.Alias != "" {
			result[identifier] = *session.Alias
		}
	}

	return result
}
