package message

func (m *Message) ValueAny(key string) any {
	if m.Values == nil {
		return nil
	}

	return m.Values[key]
}
