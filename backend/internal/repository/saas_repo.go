package repository

import (
	"gorm.io/gorm"

	"vendorsage/backend/internal/models"
)

type SaaSRepository struct {
	db *gorm.DB
}

func NewSaaSRepository(db *gorm.DB) *SaaSRepository {
	return &SaaSRepository{db: db}
}

func (r *SaaSRepository) Create(sub *models.SaaSSubscription) error {
	return r.db.Create(sub).Error
}

func (r *SaaSRepository) FindAll() ([]models.SaaSSubscription, error) {
	var subs []models.SaaSSubscription
	err := r.db.Order("created_at DESC").Find(&subs).Error
	return subs, err
}

type SignalBreakdown struct {
	SignalType string `json:"signal_type"`
	Count      int64  `json:"count"`
}

type SaaSSummary struct {
	TotalMonthly float64          `json:"total_monthly"`
	ToolCount    int64            `json:"tool_count"`
	Signals      []SignalBreakdown `json:"signals"`
}

type summaryRow struct {
	TotalMonthly float64 `gorm:"column:total_monthly"`
	ToolCount    int64   `gorm:"column:tool_count"`
}

func (r *SaaSRepository) Summary() (SaaSSummary, error) {
	var row summaryRow
	err := r.db.Model(&models.SaaSSubscription{}).
		Select("COUNT(DISTINCT product_name) as tool_count, COALESCE(SUM(CASE WHEN billing_cycle='annual' THEN estimated_cost/12 ELSE estimated_cost END), 0) as total_monthly").
		Scan(&row).Error
	if err != nil {
		return SaaSSummary{}, err
	}

	signals := make([]SignalBreakdown, 0)
	err = r.db.Model(&models.SaaSSubscription{}).
		Select("signal_type, COUNT(*) as count").
		Group("signal_type").
		Order("count DESC").
		Scan(&signals).Error
	if err != nil {
		return SaaSSummary{}, err
	}

	return SaaSSummary{
		TotalMonthly: row.TotalMonthly,
		ToolCount:    row.ToolCount,
		Signals:      signals,
	}, nil
}
