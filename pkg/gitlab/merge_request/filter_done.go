package merge_request

func FilterDone(v []*Request) []*Request {
	var result []*Request

	for _, element := range v {
		if element.Done() {
			continue
		}

		result = append(result, element)
	}

	return result
}
