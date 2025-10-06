package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	mcpadapter "github.com/i2y/langchaingo-mcp-adapter"
	"github.com/mark3labs/mcp-go/client"
	flag "github.com/spf13/pflag"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	ctx := context.Background()
	// Command-line flags
	model := flag.String("model", "google/gemini-2.5-flash-lite", "OpenRouter model to use (see https://openrouter.ai/models)")
	// prompt := flag.String("prompt", "Check the health of my database and identify any issues.", "Prompt to send to the model")
	prompt := flag.String("prompt", `What are 5 of the albums?`, "Prompt to send to the model")
	// prompt := flag.String("prompt", "explain the following query: SELECT * FROM album JOIN artist using (artistID)", "Prompt to send to the model")
	// prompt := flag.String("prompt", "execute the following query and return me the results: SELECT * FROM album JOIN artist using (artistID) limit 5", "Prompt to send to the model")
	temperature := flag.Float64("temp", 0.8, "Temperature for response generation (0.0-2.0)")
	streaming := flag.Bool("stream", true, "Use streaming mode")
	flag.Parse()

	mcpClient, err := client.NewSSEMCPClient("http://localhost:8000/sse")
	if err != nil {
		log.Fatalf("Failed to create MCP client: %v", err)
	}
	defer mcpClient.Close()
	slog.Info("initialized mcp client and server")

	err = mcpClient.Start(ctx)
	if err != nil {
		panic(err)
	}

	// Create the adapter
	adapter, err := mcpadapter.New(mcpClient)
	if err != nil {
		slog.Error("Failed to create adapter", slog.Any("error", err))
		os.Exit(1)
	}

	// Get all tools from MCP server
	mcpTools, err := adapter.Tools()
	if err != nil {
		slog.Error("failed to get tools", slog.Any("error", err))
		os.Exit(1)
	}
	for _, t := range mcpTools {
		slog.Info("retrieved tool from mcp server", slog.String("name", t.Name()), slog.String("description", t.Description()))
	}

	// OpenRouter provides access to multiple LLM providers through a unified API
	// Get your API key from https://openrouter.ai/
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		slog.Error("Please set OPENROUTER_API_KEY environment variable\nGet your key from https://openrouter.ai/")
		os.Exit(1)
	}

	fmt.Println("🚀 OpenRouter CLI - langchaingo (with mcp)")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Model: %s\n", *model)
	fmt.Printf("Temperature: %.1f\n", *temperature)
	fmt.Printf("Streaming: %v\n", *streaming)
	fmt.Printf("Prompt: %s\n", *prompt)
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

	agent := agents.NewOneShotAgent(
		llm, mcpTools, agents.WithMaxIterations(100),
	)
	executor := agents.NewExecutor(agent)

	// Use the agent
	result, err := chains.Run(
		ctx,
		executor,
		*prompt,
	)
	if err != nil {
		slog.Error("agent execution error", slog.Any("error", err))
	}

	slog.Info("Agent", slog.String("response", result))

	// Generate response
	// opts := []llms.CallOption{
	// 	llms.WithTemperature(*temperature),
	// }

	// if *streaming {
	// 	opts = append(opts, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
	// 		fmt.Print(string(chunk))
	// 		return nil
	// 	}))
	// }

	// response, err := llms.GenerateFromSinglePrompt(ctx, llm, *prompt, opts...)
	// if err != nil {
	// 	slog.Error("generating from single prompt", slog.Any("error", err))
	// }

	// if !*streaming && err == nil {
	// 	fmt.Println(response)
	// }

	// fmt.Println()

	// if err != nil {
	// 	if strings.Contains(err.Error(), "429") {
	// 		fmt.Println("⚠️  Rate limit reached. Free tier models are limited to 1 request per minute.")
	// 		fmt.Println("    Try using a different model with -model flag")
	// 	} else {
	// 		log.Printf("Error: %v\n", err)
	// 	}
	// 	os.Exit(1)
	// }

	// fmt.Println("✅ Success!")
}
