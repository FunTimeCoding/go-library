package message

func NewSlice(
	role string,
	content string,
) []*Message {
	return []*Message{New(role, content)}
}
