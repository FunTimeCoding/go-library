package issue

func GroupByStatus(input []*Issue) map[string][]*Issue {
	result := make(map[string][]*Issue)

	for _, issue := range input {
		result[issue.Status] = append(result[issue.Status], issue)
	}

	return result
}
