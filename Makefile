.PHONY: run

run:
	# google/gemini-2.5-flash-lite
	# $0.10/M input tokens, $0.40/M output tokens
	# go run . --prompt "What are some good youtube videos for learning elixir, make sure that they are highly liked"
	go run .

database-mcp:
	DB_TYPE=pg \
	DB_CONFIG="{\"host\":\"localhost\",\"port\":5432,\"user\":\"postgres\",\"password\":\"postgres\",\"dbname\":\"postgres\"}" \
	uvx database-mcp

google-flash:
	# $0.30/M input tokens, $2.50/M output tokens
	go run . --model google/gemini-2.5-flash

google-pro:
	# Starting at $1.25/M input tokens, Starting at $10/M output tokens
	go run . --model google/gemini-2.5-pro

claude-sonnet:
	# Starting at $3/M input tokens, Starting at $15/M output tokens
	go run . --model anthropic/claude-sonnet-4

gpt-5:
	# $0.625/M input tokens, $5/M output tokens, $5/K web search (until September 20)
	go run . --model openai/gpt-5

gpt-4:
	# Starting at $0.15/M input tokens, Starting at $0.60/M output tokens
	go run . --model openai/gpt-4o-mini
