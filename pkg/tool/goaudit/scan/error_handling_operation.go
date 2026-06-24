package scan

type errorHandlingOperation struct {
	OperationID string                           `yaml:"operationId"`
	Responses   map[string]errorHandlingResponse `yaml:"responses"`
}
