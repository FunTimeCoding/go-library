package resource

import "fmt"

func StripArgocdNoise(
	object map[string]interface{},
	filtered *[]string,
) {
	status, okay := object["status"].(map[string]interface{})

	if !okay {
		return
	}

	if resources, okay := status["resources"].([]interface{}); okay {
		*filtered = append(
			*filtered,
			fmt.Sprintf(
				"status.resources (%d entries, all synced)",
				len(resources),
			),
		)
		delete(status, "resources")
	}

	if history, okay := status["history"].([]interface{}); okay {
		*filtered = append(
			*filtered,
			fmt.Sprintf("status.history (%d entries)", len(history)),
		)
		delete(status, "history")
	}

	if result, okay := status["operationState"].(map[string]interface{}); okay {
		if syncResult, okay := result["syncResult"].(map[string]interface{}); okay {
			if resources, okay := syncResult["resources"].([]interface{}); okay {
				*filtered = append(
					*filtered,
					fmt.Sprintf(
						"status.operationState.syncResult.resources (%d entries)",
						len(resources),
					),
				)
				delete(syncResult, "resources")
			}
		}
	}
}
