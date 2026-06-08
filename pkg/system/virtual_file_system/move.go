package virtual_file_system

func (s *System) Move(
	from string,
	to string,
) {
	s.moves = append(s.moves, PendingMove{From: from, To: to})
}
