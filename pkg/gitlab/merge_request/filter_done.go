package merge_request

func FilterDone(v []*Request) []*Request {
	var result []*Request

	for _, e := range v {
		if e.Done() {
			continue
		}

		result = append(result, e)
	}

	return result
}
