package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"

	"vendorsage/backend/internal/models"
)

type AIResult struct {
	Expense   *models.ExpenseExtracted `json:"expense"`
	SaaS      *models.SaaSExtracted    `json:"saas"`
	Uncertain bool                     `json:"uncertain"`
}

type AIService struct {
	client anthropic.Client
}

func NewAIService(apiKey string) *AIService {
	client := anthropic.NewClient(option.WithAPIKey(apiKey))
	return &AIService{client: client}
}

func (s *AIService) ExtractFromEmail(from, subject, body, date string) (*AIResult, error) {
	result, err := s.callClaude(from, subject, body, date)
	if err != nil {
		log.Printf("[AI] first attempt failed for %q: %v — retrying", subject, err)
		result, err = s.callClaude(from, subject, body, date)
		if err != nil {
			log.Printf("[AI] retry also failed for %q: %v — marking uncertain", subject, err)
			return &AIResult{Uncertain: true}, nil
		}
	}
	return result, nil
}

func (s *AIService) callClaude(from, subject, body, date string) (*AIResult, error) {
	msg, err := s.client.Messages.New(context.Background(), anthropic.MessageNewParams{
		Model:     anthropic.ModelClaudeHaiku4_5_20251001,
		MaxTokens: 1024,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(buildPrompt(from, subject, body, date))),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("claude API error: %w", err)
	}

	raw := msg.Content[0].Text
	log.Printf("[AI] raw response for %q: %s", subject, raw)

	jsonStr := extractJSON(raw)
	var result AIResult
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, fmt.Errorf("JSON parse error: %w — raw: %s", err, raw)
	}
	return &result, nil
}

// extractJSON finds the first {...} block in the string, handling nested braces.
func extractJSON(s string) string {
	start := strings.Index(s, "{")
	if start == -1 {
		return s
	}
	depth := 0
	for i := start; i < len(s); i++ {
		switch s[i] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return s[start : i+1]
			}
		}
	}
	return s[start:]
}

func buildPrompt(from, subject, body, date string) string {
	return fmt.Sprintf(`You are a strict financial email parser. Extract only confirmed data from this email.

Email:
From: %s
Subject: %s
Date: %s
Body: %s

Return ONLY a valid JSON object. No explanation, no markdown. Use this exact shape:
{"expense":{"merchant_name":"","amount":0.00,"currency":"USD","category":"food_delivery|travel|software|shopping|utilities|entertainment|other","transaction_date":"YYYY-MM-DD"},"saas":{"product_name":"","signal_type":"welcome|invoice|renewal|trial_expiring|usage_report","billing_cycle":"monthly|annual|unknown","estimated_cost":0.00},"uncertain":false}

Expense rules (set to null if none of these apply):
- Only extract if money was ACTUALLY CHARGED. Words like "charged", "payment received", "billed", "your receipt", "order total" confirm a charge.
- DO NOT extract if the email only mentions a price as an upsell, upgrade suggestion, or future offer.
- DO NOT extract if the email is about a trial that has not yet converted to a paid plan.

SaaS rules (set to null if not a SaaS tool):
- signal_type must be one of: welcome, invoice, renewal, trial_expiring, usage_report.
- For trial_expiring emails: set estimated_cost to 0.00 since no charge has occurred yet.
- For free plan welcome emails: set estimated_cost to 0.00.
- estimated_cost should only reflect what was ACTUALLY CHARGED, not a potential future price.

SaaS product_name rule: Use the company name derived from the sender's email domain as a fallback if the product name is not explicitly stated in the body. For example: "no-reply@netflix.com" → "Netflix", "billing@github.com" → "GitHub".

Both expense and saas can be non-null for the same email (e.g. a SaaS invoice is both a charge and a subscription signal).
Set uncertain to true only if the email is completely unrelated to finances or subscriptions.`,
		from, subject, date, body)
}
