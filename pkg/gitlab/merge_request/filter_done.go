package merge_request

func FilterDone(input []*Request) []*Request {
	var result []*Request

	for _, element := range input {
		if element.Done() {
			continue
		}

		result = append(result, element)
	}

	return result
}
