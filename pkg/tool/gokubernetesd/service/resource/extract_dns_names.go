package resource

func ExtractDnsNames(object map[string]interface{}) []string {
	raw, okay := object["spec"].(map[string]interface{})["dnsNames"]

	if !okay {
		return nil
	}

	list, okay := raw.([]interface{})

	if !okay {
		return nil
	}

	var result []string

	for _, item := range list {
		if s, okay := item.(string); okay {
			result = append(result, s)
		}
	}

	return result
}
