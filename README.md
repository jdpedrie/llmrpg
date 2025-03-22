# LLMRPG

A role-playing game system built with Go, PostgreSQL, and pgvector.

## Project Setup

This project has been converted from EdgeDB to PostgreSQL with pgvector for vector search capabilities.

### Prerequisites

- Go 1.24+
- Docker and Docker Compose
- SQLC (for development)
- Goose (for database migrations)

## Getting Started

1. Start the PostgreSQL database:

```bash
docker-compose up -d postgres
```

2. Apply database migrations:

```bash
goose -dir migrations postgres "postgresql://postgres:postgres@localhost:5432/llmrpg?sslmode=disable" up
```

3. Run the application:

```bash
go run ./cmd/main.go
```

To start the gRPC server:

```bash
go run ./cmd/main.go serve
```

## Database Management

### Migrations

Migrations are managed using Goose. To create a new migration:

```bash
goose -dir migrations create migration_name sql
```

To apply migrations:

```bash
goose -dir migrations postgres "postgresql://postgres:postgres@localhost:5432/llmrpg?sslmode=disable" up
```

To roll back a migration:

```bash
goose -dir migrations postgres "postgresql://postgres:postgres@localhost:5432/llmrpg?sslmode=disable" down
```

### SQLC

The project uses SQLC to generate typesafe Go code from SQL queries.

To regenerate the SQL code:

```bash
sqlc generate
```

## Development

The main components of the system are:

- `model`: Core domain models and proto conversions
- `pkg/postgres`: SQLC-generated database access and utility functions
- `game`: Game management and engine functionality
- `cmd`: Command line interfaces and server implementation

## Vector Search

The `history` table uses pgvector to store embeddings, allowing semantic search over game history entries:

```go
// Example of searching similar history entries
history, err := db.SearchSimilarHistory(ctx, SearchSimilarHistoryParams{
    Embedding: embedding, // vector(1536)
    GameID:    gameID,
    Limit:     10,
})
```