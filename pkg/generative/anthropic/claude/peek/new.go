package peek

func New(messages []string) *Peek {
	return &Peek{UserMessages: messages}
}
