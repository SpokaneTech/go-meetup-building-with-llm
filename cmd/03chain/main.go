package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/memory"
)

func main() {
	// Command-line flags
	model := flag.String("model", "google/gemini-2.5-flash-lite", "OpenRouter model to use (see https://openrouter.ai/models)")
	flag.Parse()

	// OpenRouter provides access to multiple LLM providers through a unified API
	// Get your API key from https://openrouter.ai/
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		slog.Error("Please set OPENROUTER_API_KEY environment variable\nGet your key from https://openrouter.ai/")
		os.Exit(1)
	}

	fmt.Println("🚀 OpenRouter CLI - langchaingo")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Model: %s\n", *model)
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println()

	// Create an OpenAI-compatible client configured for OpenRouter
	llm, err := openai.New(
		openai.WithModel(*model),
		openai.WithBaseURL("https://openrouter.ai/api/v1"),
		openai.WithToken(apiKey),
	)
	if err != nil {
		slog.Error("unable to create openai compatible llm", slog.Any("error", err))
	}

	// Create conversation memory
	chatMemory := memory.NewConversationBuffer()

	// Create conversation chain
	// The built-in conversation chain includes a default prompt template
	// and handles memory automatically
	conversationChain := chains.NewConversation(llm, chatMemory)

	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Advanced Chat Application (type 'quit' to exit)")
	fmt.Println("----------------------------------------")

	for {
		fmt.Print("You: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			break
		}

		// Run the chain with the input
		result, err := chains.Run(ctx, conversationChain, input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Printf("AI: %s\n\n", result)
	}

	fmt.Println("Goodbye!")
}
