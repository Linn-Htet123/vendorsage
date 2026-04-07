package models

import "time"

type ExpenseCategory string

const (
	CategoryFoodDelivery ExpenseCategory = "food_delivery"
	CategoryTravel       ExpenseCategory = "travel"
	CategorySoftware     ExpenseCategory = "software"
	CategoryShopping     ExpenseCategory = "shopping"
	CategoryUtilities    ExpenseCategory = "utilities"
	CategoryEntertainment ExpenseCategory = "entertainment"
	CategoryOther        ExpenseCategory = "other"
)

type Expense struct {
	ID              uint            `gorm:"primaryKey" json:"id"`
	EmailID         uint            `gorm:"not null;index" json:"email_id"`
	MerchantName    string          `json:"merchant_name"`
	Amount          float64         `json:"amount"`
	Currency        string          `gorm:"default:'USD'" json:"currency"`
	Category        ExpenseCategory `gorm:"index" json:"category"`
	TransactionDate time.Time       `gorm:"index" json:"transaction_date"`
	Uncertain       bool            `gorm:"default:false" json:"uncertain"`
	CreatedAt       time.Time       `json:"created_at"`
}

type ExpenseExtracted struct {
	MerchantName    string  `json:"merchant_name"`
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
	Category        string  `json:"category"`
	TransactionDate string  `json:"transaction_date"`
}
