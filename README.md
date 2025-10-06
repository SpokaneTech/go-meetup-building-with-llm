# Building with an LLM

[Langchain Go](https://github.com/tmc/langchaingo)
[Langchain Go MCP adapter](https://github.com/i2y/langchaingo-mcp-adapter)
[Postgres MCP server](https://github.com/TheRaLabs/legion-mcp)
[Postgres MCP server](https://github.com/vinsidious/mcp-pg-schema)
[Neon postgres MCP server](https://github.com/neondatabase/mcp-server-neon)
[DBs](https://github.com/neondatabase-labs/postgres-sample-dbs)

## RAG

[Example RAG](https://github.com/build-on-aws/rag-golang-postgresql-langchain)

## Steps

1. Setup makefile to switch different models
2. setup docker/podman with postgres and sample data
3. setup local mcp to access postgres data
4. Configure langchain go to access mcp server
5. Install rdbc-mcp (cargo install --git https://github.com/rbatis/rbdc-mcp.git)
