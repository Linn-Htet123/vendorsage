# Frontend

React dashboard for viewing extracted expenses and SaaS subscriptions, with a built-in email upload tool.

## Setup

```bash
npm install
```

Optionally copy `.env.example` if the backend runs on a different port:

```bash
cp .env.example .env
```

## Run

```bash
npm run dev
```

Opens at http://localhost:5173. Requires the backend to be running at http://localhost:8080.

## Commands

```bash
npm run dev        # start dev server
npm run build      # production build (output: dist/)
npm run preview    # preview production build
npx tsc --noEmit   # type check
```

## Pages

| Route | Description |
|-------|-------------|
| `/spending` | Expense table with category + date filters and spend summary |
| `/saas` | SaaS subscription list with monthly cost summary and email upload |

## Uploading emails

On the **SaaS** page, click **Upload Emails (JSON)** and select a `.json` file:

```json
[
  {
    "from": "billing@github.com",
    "to": "you@example.com",
    "subject": "Your GitHub invoice",
    "body": "Your monthly plan of $4.00 has been charged.",
    "date": "2025-04-01"
  }
]
```

## Environment variables

| Variable | Default | Description |
|----------|---------|-------------|
| `VITE_API_URL` | `http://localhost:8080` | Backend base URL |
