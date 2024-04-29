package system

func LookupFirst(address string) string {
	results := Lookup(address)

	if len(results) > 0 {
		return results[0]
	}

	return ""
}
