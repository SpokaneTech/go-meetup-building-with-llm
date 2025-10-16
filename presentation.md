---
title: "LangchainGo"
author: Brennon Loveless
theme:
  name: catppuccin-macchiato
---

What is LangchainGo?
===

<!-- new_lines: 2 -->

Interacting with LLMs generally mean chaining together a bunch of different API calls. Building up a context.
And then asking questions.

The things you expect from a conversation don't come for free. For example, having a conversation with an LLM
requires that the context of the conversation is passed back into every LLM call.

```
Demo: go run cmd/01chat/main.go
You: Tell me a story about a Gnome?
AI: <a story>
You: What was that story about?
AI: <no idea>
```

<!-- end_slide -->

Using the library for conversational history
===

So how do we add conversational history. The good news is that the library provides a way to accomplish just that.

```
Demo: go run cmd/02memory/main.go
You: Tell me a short story about a robot?
AI: <short story about a robot>
You: Continue the story and include something about a golden record.
AI: <continues the story using the context from the previous conversation>
````

<!-- end_slide -->

What is a chain?
===

So far we've just been using pieces of the library, so what makes this a chain?

Chain create multi step applications that can gather information from third party APIs or sources or other LLMs.

This is actually the same demo as above but the chains generally provide memory as well.

Other types of [chains](https://tmc.github.io/langchaingo/docs/modules/chains/)
Chaining [chains](https://github.com/tmc/langchaingo/blob/main/examples/sequential-chain-example/sequential_chain_example.go)

```
Demo: go run cmd/03chain/main.go
You: Tell me a short story about a robot?
AI: <short story about a robot>
You: Continue the story and include something about a golden record.
AI: <continues the story using the context from the previous conversation>
```

<!-- end_slide -->

How to use external functionality with the LLM?
===

External deterministic functionality is provided to the LLM through tool calls via an MCP server.

In this example we'll provide the LLM access to a local postgres database server and ask it questions.

```
Setup: podman compose up -d
Demo: go run cmd/04toolcalling/main.go
```

So the LLM was able to use the MCP server to inspect the database and figure out how to query for 5 of
the albums from the postgres database.

<!-- end_slide -->

Other resources
===

[Examples](https://github.com/tmc/langchaingo/tree/main/examples)

[RAG](https://github.com/build-on-aws/rag-golang-postgresql-langchain)

[Ollama and Redis example](https://github.com/tmc/langchaingo/blob/main/examples/redis-vectorstore-example/redis_vectorstore_example.go)

[Openrouter](https://openrouter.ai/)

[Openrouter Gemini](https://openrouter.ai/google/gemini-2.5-flash-lite)

[Openrouter Claude Sonnet](https://openrouter.ai/anthropic/claude-sonnet-4.5)

[Openrouter Claude Haiku](https://openrouter.ai/anthropic/claude-haiku-4.5)
