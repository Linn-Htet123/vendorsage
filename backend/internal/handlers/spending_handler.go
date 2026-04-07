package handlers

import (
	"time"

	"github.com/gin-gonic/gin"

	"vendorsage/backend/internal/repository"
	"vendorsage/backend/pkg/utils"
)

type SpendingHandler struct {
	repo *repository.ExpenseRepository
}

func NewSpendingHandler(repo *repository.ExpenseRepository) *SpendingHandler {
	return &SpendingHandler{repo: repo}
}

// List handles GET /api/spending
func (h *SpendingHandler) List(c *gin.Context) {
	filter := repository.ExpenseFilter{
		Category: c.Query("category"),
	}

	if s := c.Query("date_start"); s != "" {
		t, err := time.Parse("2006-01-02", s)
		if err != nil {
			utils.BadRequest(c, "invalid date_start format, use YYYY-MM-DD")
			return
		}
		filter.DateStart = &t
	}

	if s := c.Query("date_end"); s != "" {
		t, err := time.Parse("2006-01-02", s)
		if err != nil {
			utils.BadRequest(c, "invalid date_end format, use YYYY-MM-DD")
			return
		}
		filter.DateEnd = &t
	}

	expenses, err := h.repo.FindAll(filter)
	if err != nil {
		utils.InternalError(c, "failed to fetch expenses")
		return
	}
	if expenses == nil {
		expenses = []repository.ExpenseWithSource{}
	}
	utils.OK(c, expenses)
}

// Summary handles GET /api/spending/summary
func (h *SpendingHandler) Summary(c *gin.Context) {
	total, breakdown, err := h.repo.Summary()
	if err != nil {
		utils.InternalError(c, "failed to fetch summary")
		return
	}
	utils.OK(c, gin.H{
		"total":     total,
		"breakdown": breakdown,
	})
}
