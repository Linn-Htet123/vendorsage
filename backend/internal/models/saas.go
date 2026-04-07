package models

import "time"

type SignalType string
type BillingCycle string

const (
	SignalWelcome        SignalType = "welcome"
	SignalInvoice        SignalType = "invoice"
	SignalRenewal        SignalType = "renewal"
	SignalTrialExpiring  SignalType = "trial_expiring"
	SignalUsageReport    SignalType = "usage_report"
)

const (
	BillingMonthly BillingCycle = "monthly"
	BillingAnnual  BillingCycle = "annual"
	BillingUnknown BillingCycle = "unknown"
)

type SaaSSubscription struct {
	ID            uint         `gorm:"primaryKey" json:"id"`
	EmailID       uint         `gorm:"not null;index" json:"email_id"`
	ProductName   string       `gorm:"index" json:"product_name"`
	SignalType    SignalType   `json:"signal_type"`
	BillingCycle  BillingCycle `json:"billing_cycle"`
	EstimatedCost float64      `json:"estimated_cost"`
	Uncertain     bool         `gorm:"default:false" json:"uncertain"`
	CreatedAt     time.Time    `json:"created_at"`
}

func (SaaSSubscription) TableName() string { return "saas_subscriptions" }

type SaaSExtracted struct {
	ProductName   string  `json:"product_name"`
	SignalType    string  `json:"signal_type"`
	BillingCycle  string  `json:"billing_cycle"`
	EstimatedCost float64 `json:"estimated_cost"`
}
