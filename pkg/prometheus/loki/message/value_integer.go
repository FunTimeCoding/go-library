package message

func (m *Message) ValueInteger(key string) int {
	if m.Values == nil {
		return 0
	}

	if v, okay := m.Values[key]; okay {
		return int(v.(float64))
	}

	return 0
}
