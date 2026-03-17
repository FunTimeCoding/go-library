package message

func New(
	role string,
	content string,
) *Message {
	return &Message{Role: role, Content: content}
}
