# Backend

Go REST API that ingests emails, extracts expenses and SaaS subscriptions using Claude AI, and stores results in PostgreSQL.

## Setup

```bash
cp .env.example .env
```

Edit `.env` and set your `ANTHROPIC_API_KEY`. Make sure PostgreSQL is running (`docker compose up -d` from the root).

## Run

**With hot reload (recommended for development):**

```bash
# Install air once
go install github.com/air-verse/air@latest

# Start with auto-restart on file save
air
```

**Without hot reload:**

```bash
go run ./cmd/main.go
```

Server starts at `http://localhost:8080`.

> Note: Go does not hot-reload. If you use `go run`, you must restart manually after every code change. Use `air` to avoid this.

## Other commands

```bash
go build ./...    # compile
go test ./...     # run all tests
go vet ./...      # lint
```

## API

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/emails/upload` | Upload emails for AI processing |
| GET | `/api/spending` | List expenses (`?category=&date_start=&date_end=`) |
| GET | `/api/spending/summary` | Total spend grouped by category |
| GET | `/api/saas` | List SaaS subscriptions |
| GET | `/api/saas/summary` | Monthly SaaS spend + tool count |

## Upload format

```json
{
  "emails": [
    {
      "from": "receipt@uber.com",
      "to": "you@example.com",
      "subject": "Your Uber Eats receipt",
      "body": "You were charged $24.50 on April 5, 2025.",
      "date": "2025-04-05"
    }
  ]
}
```

## Environment variables

| Variable | Default | Description |
|----------|---------|-------------|
| `ANTHROPIC_API_KEY` | — | **Required.** Your Anthropic API key |
| `DB_HOST` | `localhost` | PostgreSQL host |
| `DB_PORT` | `5433` | PostgreSQL port (mapped via Docker) |
| `DB_USER` | `postgres` | PostgreSQL user |
| `DB_PASSWORD` | `postgres` | PostgreSQL password |
| `DB_NAME` | `vendorsage` | Database name |
| `SERVER_PORT` | `8080` | HTTP server port |
