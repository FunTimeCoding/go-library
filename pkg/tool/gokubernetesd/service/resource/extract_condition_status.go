package resource

func ExtractConditionStatus(object map[string]interface{}) string {
	status, okay := object["status"].(map[string]interface{})

	if !okay {
		return "Unknown"
	}

	conditions, okay := status["conditions"].([]interface{})

	if !okay {
		return "Unknown"
	}

	for _, raw := range conditions {
		condition, okay := raw.(map[string]interface{})

		if !okay {
			continue
		}

		if condition["type"] == "Ready" {
			if condition["status"] == "True" {
				return "Ready"
			}

			reason, _ := condition["reason"].(string)

			if reason != "" {
				return reason
			}

			return "NotReady"
		}
	}

	return "Unknown"
}
