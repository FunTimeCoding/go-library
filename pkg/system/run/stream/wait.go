package stream

func (s *Stream) Wait() error {
	return s.command.Wait()
}
