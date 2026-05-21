package web

func isEditable(kind string) bool {
	return kind == "complete" || kind == "summarize" || kind == "moment"
}
