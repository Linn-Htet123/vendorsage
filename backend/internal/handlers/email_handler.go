package handlers

import (
	"github.com/gin-gonic/gin"

	"vendorsage/backend/internal/models"
	"vendorsage/backend/internal/services"
	"vendorsage/backend/pkg/utils"
)

type EmailHandler struct {
	processor *services.EmailProcessor
}

func NewEmailHandler(processor *services.EmailProcessor) *EmailHandler {
	return &EmailHandler{processor: processor}
}

type uploadRequest struct {
	Emails []models.EmailInput `json:"emails" binding:"required,min=1"`
}

type uploadResponse struct {
	Processed    int `json:"processed"`
	Skipped      int `json:"skipped"`
	Failed       int `json:"failed"`
	Total        int `json:"total"`
	Expense      int `json:"expense"`
	SaaS         int `json:"saas"`
	Both         int `json:"both"`
	NonFinancial int `json:"non_financial"`
	Uncertain    int `json:"uncertain"`
}

// Upload handles POST /api/emails/upload
func (h *EmailHandler) Upload(c *gin.Context) {
	var req uploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "invalid request body: "+err.Error())
		return
	}

	r := h.processor.ProcessBatch(req.Emails)

	utils.Created(c, uploadResponse{
		Processed:    r.Processed,
		Skipped:      r.Skipped,
		Failed:       r.Failed,
		Total:        len(req.Emails),
		Expense:      r.Expense,
		SaaS:         r.SaaS,
		Both:         r.Both,
		NonFinancial: r.NonFinancial,
		Uncertain:    r.Uncertain,
	})
}
