package constant

const (
	HostEnvironment string = "OLLAMA_HOST"

	Port int = 11434

	Llama31   = "llama3.1"    // 8b
	Llama32   = "llama3.2"    // 3b
	Llama321b = "llama3.2:1b" // 1b

	SystemRole    = "system"
	UserRole      = "user"
	AssistantRole = "assistant"

	NotationFormat = "json"
)
