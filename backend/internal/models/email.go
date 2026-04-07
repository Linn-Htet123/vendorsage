package models

import "time"

type EmailType string

const (
	EmailTypeExpense      EmailType = "expense"
	EmailTypeSaaS         EmailType = "saas"
	EmailTypeBoth         EmailType = "both"
	EmailTypeNonFinancial EmailType = "non_financial"
	EmailTypeUncertain    EmailType = "uncertain"
)

type Email struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	From        string    `gorm:"not null" json:"from"`
	To          string    `json:"to"`
	Subject     string    `json:"subject"`
	Body        string    `gorm:"type:text" json:"body"`
	Date        time.Time `json:"date"`
	ContentHash string    `gorm:"uniqueIndex;size:64" json:"content_hash"`
	Type        EmailType `gorm:"default:'non_financial'" json:"type"`
	CreatedAt   time.Time `json:"created_at"`

	Expenses          []Expense          `gorm:"foreignKey:EmailID" json:"expenses,omitempty"`
	SaaSSubscriptions []SaaSSubscription `gorm:"foreignKey:EmailID" json:"saas_subscriptions,omitempty"`
}

// EmailInput is used for JSON/CSV upload parsing.
type EmailInput struct {
	From    string `json:"from" binding:"required"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Date    string `json:"date"`
}
