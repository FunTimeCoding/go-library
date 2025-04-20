package chat_response

type Response struct {
	Text             string
	Total            int64
	Load             int64
	PromptEvaluation int64
	Evaluation       int64
	PromptTokens     float64
	Tokens           float64
}
