package main

import (
	anthropic "github.com/funtimecoding/go-library/pkg/generative/anthropic/example"
	gguf "github.com/funtimecoding/go-library/pkg/generative/gguf/example"
	langchain "github.com/funtimecoding/go-library/pkg/generative/langchain/example"
	modelContext "github.com/funtimecoding/go-library/pkg/generative/model_context/example"
	ollama "github.com/funtimecoding/go-library/pkg/generative/ollama/example"
	openWebUI "github.com/funtimecoding/go-library/pkg/generative/open_webui/example"
	openai "github.com/funtimecoding/go-library/pkg/generative/openai/example"
)

func main() {
	if false {
		gguf.Read()
	}

	if false {
		ollama.Chat()
		ollama.Benchmark()
		ollama.ClassifyAlert()
		ollama.Fast()
		ollama.Stream()
		ollama.Simple()
		ollama.Show()
		ollama.List()
		ollama.Embed()
		ollama.Running()
		ollama.Heartbeat()
	}

	if false {
		openai.Token()
		openai.Official()
		openai.Alternate()
	}

	if false {
		anthropic.Official()
		anthropic.Alternate()
	}

	if false {
		langchain.Chroma()
		langchain.Function()
		langchain.Local()
	}

	if false {
		openWebUI.Load()
	}

	if false {
		modelContext.Server()
	}
}
