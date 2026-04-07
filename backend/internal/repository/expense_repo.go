package repository

import (
	"time"

	"gorm.io/gorm"

	"vendorsage/backend/internal/models"
)

type ExpenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) *ExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (r *ExpenseRepository) Create(expense *models.Expense) error {
	return r.db.Create(expense).Error
}

type ExpenseFilter struct {
	Category  string
	DateStart *time.Time
	DateEnd   *time.Time
}

// ExpenseWithSource extends Expense with the source email fields for traceability.
type ExpenseWithSource struct {
	models.Expense
	SourceSubject string `json:"source_subject"`
	SourceDate    string `json:"source_date"`
}

func (r *ExpenseRepository) FindAll(f ExpenseFilter) ([]ExpenseWithSource, error) {
	q := r.db.Model(&models.Expense{}).
		Select("expenses.*, emails.subject as source_subject, TO_CHAR(emails.date, 'YYYY-MM-DD') as source_date").
		Joins("LEFT JOIN emails ON emails.id = expenses.email_id")

	if f.Category != "" {
		q = q.Where("expenses.category = ?", f.Category)
	}
	if f.DateStart != nil {
		q = q.Where("expenses.transaction_date >= ?", f.DateStart)
	}
	if f.DateEnd != nil {
		q = q.Where("expenses.transaction_date <= ?", f.DateEnd)
	}

	var expenses []ExpenseWithSource
	err := q.Order("expenses.transaction_date DESC").Find(&expenses).Error
	return expenses, err
}

type CategorySummary struct {
	Category string  `json:"category"`
	Total    float64 `json:"total"`
	Count    int64   `json:"count"`
}

func (r *ExpenseRepository) Summary() (float64, []CategorySummary, error) {
	var total float64
	r.db.Model(&models.Expense{}).Select("COALESCE(SUM(amount), 0)").Scan(&total)

	breakdown := make([]CategorySummary, 0)
	err := r.db.Model(&models.Expense{}).
		Select("category, SUM(amount) as total, COUNT(*) as count").
		Group("category").
		Order("total DESC").
		Scan(&breakdown).Error

	return total, breakdown, err
}
