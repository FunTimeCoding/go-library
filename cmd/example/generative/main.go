package main

import (
	anthropic "github.com/funtimecoding/go-library/pkg/generative/anthropic/example"
	gguf "github.com/funtimecoding/go-library/pkg/generative/gguf/example"
	langchain "github.com/funtimecoding/go-library/pkg/generative/langchain/example"
	mistral "github.com/funtimecoding/go-library/pkg/generative/mistral/example"
	anthropicServer "github.com/funtimecoding/go-library/pkg/generative/model_context/example/anthropic"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/example/mark"
	ollama "github.com/funtimecoding/go-library/pkg/generative/ollama/example"
	openWebUI "github.com/funtimecoding/go-library/pkg/generative/open_webui/example"
	openai "github.com/funtimecoding/go-library/pkg/generative/openai/example"
)

func main() {
	mistral.Prompt()

	if false {
		ollama.Chat()
		ollama.Benchmark()
		ollama.ClassifyAlert()
		ollama.Fast()
		ollama.Custom()
		ollama.Stream()
		ollama.Simple()
		ollama.Show()
		ollama.List()
		ollama.Embed()
		ollama.Running()
		ollama.Heartbeat()

		openai.Token()
		openai.Official()
		openai.Alternate()
		openai.Client()

		anthropic.Official()
		anthropic.Alternate()

		langchain.Chroma()
		langchain.Function()
		langchain.Local()

		openWebUI.Load()

		anthropicServer.Run()
		mark.Main()

		gguf.Read()
	}
}
