# VendorSage

VendorSage reads your emails and finds two things: money you spent, and SaaS tools you are paying for. You upload a JSON or CSV file, the app sends each email to Claude AI, and the results show up in a dashboard.

**Stack:** Go + Gin · PostgreSQL · React + Vite + shadcn/ui · Anthropic Claude Haiku

---

## Setup

**You need:** Docker and an Anthropic API key.

```bash
cp backend/.env.example backend/.env
# Open backend/.env and add your ANTHROPIC_API_KEY

docker compose up --build
```

| Service | URL |
|---|---|
| Dashboard | http://localhost:4173 |
| API | http://localhost:8080 |
| Adminer (DB) | http://localhost:8081 |

Open the dashboard and click **Import JSON / CSV** to upload `sample_emails.json`.

---

## API

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/api/emails/upload` | Upload emails (JSON or CSV) |
| `GET` | `/api/spending` | List expenses — filters: `category`, `date_start`, `date_end` |
| `GET` | `/api/spending/summary` | Total spend + breakdown by category |
| `GET` | `/api/saas` | List SaaS subscriptions |
| `GET` | `/api/saas/summary` | Monthly SaaS total + tool count |

---

## Design Decisions

**The app sends one email to Claude at a time.**
Claude returns the expense and SaaS data as JSON. The app parses that JSON and saves it to the database. If parsing fails, the app retries once. If it fails again, the email is marked uncertain and the batch continues.

**The app saves the email before it calls Claude.**
This way, if Claude fails, the email is not lost. It stays in the database. Duplicate emails are skipped using a SHA-256 hash of the content.

**The AI prompt is strict about what counts as a charge.**
The app only extracts money if the email says something like "charged", "payment received", or "billed". It ignores trial emails, upgrade suggestions, and future invoices. If the email does not name the product clearly, the app reads the sender domain instead — for example, `billing@github.com` becomes GitHub.

**The app checks AI output before saving it.**
Claude sometimes returns values that do not match the expected categories. The app checks each value and replaces anything invalid with a safe default. Nothing bad gets written to the database.

**Each expense shows which email it came from.**
The spending API joins the emails table and returns the email subject and date with every expense record.

---

## What I'd Add Next

- **Async processing** — run AI in the background so large uploads do not block the request
- **Confidence score** — a number from 0 to 1 instead of a simple uncertain flag
- **Gmail integration** — connect a real inbox with OAuth instead of uploading a file
- **Reprocess endpoint** — re-run Claude on a single email after improving the prompt
- **Tests** — unit tests for the prompt builder, JSON parser, and email classifier

---

## Time Spent

~4–5 hours · 1.5h backend · 1h database + Docker · 1h frontend · 1h debugging
