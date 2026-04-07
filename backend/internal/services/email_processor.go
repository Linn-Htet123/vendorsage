package services

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"time"

	"vendorsage/backend/internal/models"
	"vendorsage/backend/internal/repository"
)

type EmailProcessor struct {
	aiService   *AIService
	emailRepo   *repository.EmailRepository
	expenseRepo *repository.ExpenseRepository
	saasRepo    *repository.SaaSRepository
}

func NewEmailProcessor(
	ai *AIService,
	emailRepo *repository.EmailRepository,
	expenseRepo *repository.ExpenseRepository,
	saasRepo *repository.SaaSRepository,
) *EmailProcessor {
	return &EmailProcessor{ai, emailRepo, expenseRepo, saasRepo}
}

type BatchResult struct {
	Processed    int
	Skipped      int
	Failed       int
	Expense      int
	SaaS         int
	Both         int
	NonFinancial int
	Uncertain    int
}

func (p *EmailProcessor) ProcessBatch(inputs []models.EmailInput) BatchResult {
	var r BatchResult
	for _, input := range inputs {
		emailType, err := p.processOne(input)
		switch {
		case errors.Is(err, repository.ErrDuplicate):
			log.Printf("[processor] skipping duplicate: %q", input.Subject)
			r.Skipped++
		case err != nil:
			log.Printf("[processor] failed %q: %v", input.Subject, err)
			r.Failed++
		default:
			r.Processed++
			switch emailType {
			case models.EmailTypeExpense:
				r.Expense++
			case models.EmailTypeSaaS:
				r.SaaS++
			case models.EmailTypeBoth:
				r.Both++
			case models.EmailTypeUncertain:
				r.Uncertain++
			default:
				r.NonFinancial++
			}
		}
	}
	return r
}

func contentHash(input models.EmailInput) string {
	h := sha256.Sum256([]byte(input.From + input.Subject + input.Date + input.Body))
	return fmt.Sprintf("%x", h)
}

func (p *EmailProcessor) processOne(input models.EmailInput) (models.EmailType, error) {
	date, _ := time.Parse("2006-01-02", input.Date)

	email := &models.Email{
		From:        input.From,
		To:          input.To,
		Subject:     input.Subject,
		Body:        input.Body,
		Date:        date,
		ContentHash: contentHash(input),
		Type:        models.EmailTypeNonFinancial,
	}
	if err := p.emailRepo.Create(email); err != nil {
		return "", err
	}

	result, err := p.aiService.ExtractFromEmail(input.From, input.Subject, input.Body, input.Date)
	if err != nil || result == nil {
		log.Printf("[processor] AI extraction failed for %q: %v", input.Subject, err)
		p.emailRepo.UpdateType(email.ID, models.EmailTypeUncertain)
		return models.EmailTypeUncertain, nil
	}

	emailType := classifyEmail(result)
	p.emailRepo.UpdateType(email.ID, emailType)

	log.Printf("[processor] %q → type=%s", input.Subject, emailType)

	if result.Expense != nil {
		p.saveExpense(email.ID, result.Expense, result.Uncertain)
	}
	if result.SaaS != nil {
		p.saveSaaS(email.ID, result.SaaS, result.Uncertain)
	}
	return emailType, nil
}

func classifyEmail(result *AIResult) models.EmailType {
	hasExpense := result.Expense != nil
	hasSaaS := result.SaaS != nil
	switch {
	case result.Uncertain:
		return models.EmailTypeUncertain
	case hasExpense && hasSaaS:
		return models.EmailTypeBoth
	case hasExpense:
		return models.EmailTypeExpense
	case hasSaaS:
		return models.EmailTypeSaaS
	default:
		return models.EmailTypeNonFinancial
	}
}

var validCategories = map[string]bool{
	"food_delivery": true, "travel": true, "software": true,
	"shopping": true, "utilities": true, "entertainment": true, "other": true,
}

var validSignals = map[string]bool{
	"welcome": true, "invoice": true, "renewal": true,
	"trial_expiring": true, "usage_report": true,
}

func (p *EmailProcessor) saveExpense(emailID uint, e *models.ExpenseExtracted, uncertain bool) {
	category := e.Category
	if !validCategories[category] {
		log.Printf("[processor] unknown category %q — falling back to 'other'", category)
		category = "other"
	}

	txDate, err := time.Parse("2006-01-02", e.TransactionDate)
	if err != nil {
		log.Printf("[processor] unparseable transaction_date %q — using zero time", e.TransactionDate)
	}

	expense := &models.Expense{
		EmailID:         emailID,
		MerchantName:    e.MerchantName,
		Amount:          e.Amount,
		Currency:        e.Currency,
		Category:        models.ExpenseCategory(category),
		TransactionDate: txDate,
		Uncertain:       uncertain,
	}
	if err := p.expenseRepo.Create(expense); err != nil {
		log.Printf("failed to save expense: %v", err)
	}
}

func (p *EmailProcessor) saveSaaS(emailID uint, s *models.SaaSExtracted, uncertain bool) {
	signal := s.SignalType
	if !validSignals[signal] {
		log.Printf("[processor] unknown signal_type %q — falling back to 'usage_report'", signal)
		signal = "usage_report"
	}

	billing := s.BillingCycle
	if billing != "monthly" && billing != "annual" {
		billing = "unknown"
	}

	sub := &models.SaaSSubscription{
		EmailID:       emailID,
		ProductName:   s.ProductName,
		SignalType:    models.SignalType(signal),
		BillingCycle:  models.BillingCycle(billing),
		EstimatedCost: s.EstimatedCost,
		Uncertain:     uncertain,
	}
	if err := p.saasRepo.Create(sub); err != nil {
		log.Printf("failed to save saas subscription: %v", err)
	}
}
