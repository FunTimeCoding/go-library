package maps

import (
	"encoding/base64"
	"fmt"
)

func ToMarkup(m map[string]string) string {
	result := ""

	for k, v := range m {
		result += fmt.Sprintf(
			"  %s: %s\n",
			k,
			base64.StdEncoding.EncodeToString([]byte(v)),
		)
	}

	return result
}
