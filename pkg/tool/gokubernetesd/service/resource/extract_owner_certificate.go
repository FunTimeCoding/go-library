package resource

func ExtractOwnerCertificate(object map[string]interface{}) string {
	metadata, okay := object["metadata"].(map[string]interface{})

	if !okay {
		return ""
	}

	owners, okay := metadata["ownerReferences"].([]interface{})

	if !okay {
		return ""
	}

	for _, raw := range owners {
		e, okay := raw.(map[string]interface{})

		if !okay {
			continue
		}

		if e["kind"] == "Certificate" {
			name, _ := e["name"].(string)

			return name
		}
	}

	return ""
}
