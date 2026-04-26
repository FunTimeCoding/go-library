package mock_worker

func New(
	tag string,
	log *[]string,
) *MockWorker {
	return &MockWorker{Tag: tag, Log: log}
}
