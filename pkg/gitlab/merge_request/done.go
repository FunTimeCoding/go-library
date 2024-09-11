package merge_request

func (m *Request) Done() bool {
	if m.State == "merged" {
		return true
	}

	if m.State == "closed" {
		return true
	}

	return false
}
