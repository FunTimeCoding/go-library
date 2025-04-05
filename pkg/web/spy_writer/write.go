package spy_writer

func (s *Writer) Write(bytes []byte) (int, error) {
	s.Written = bytes

	return len(bytes), nil
}
