package service

func (s *Service) MemoryActivity(since string) []MemoryActivityEntry {
	versions := s.memory.VersionsSince(since, 50)
	var result []MemoryActivityEntry

	for _, v := range versions {
		result = append(
			result,
			MemoryActivityEntry{
				MemoryIdentifier: v.MemoryIdentifier,
				Name:             v.Name,
				ChangeType:       v.ChangeType,
				Source:           v.Source,
			},
		)
	}

	return result
}
