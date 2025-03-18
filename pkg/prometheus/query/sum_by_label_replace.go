package query

import "fmt"

func SumByLabelReplace(
	query string,
	matchFieldName string,
	newFieldName string,
	expression string,
) string {
	return fmt.Sprintf(
		`sum by (%s) (label_replace(%s, "%s", "$1", "%s", "%s"))`,
		newFieldName,
		query,
		newFieldName,
		matchFieldName,
		expression,
	)
}
