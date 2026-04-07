package handlers

import (
	"github.com/gin-gonic/gin"

	"vendorsage/backend/internal/repository"
	"vendorsage/backend/pkg/utils"
)

type SaaSHandler struct {
	repo *repository.SaaSRepository
}

func NewSaaSHandler(repo *repository.SaaSRepository) *SaaSHandler {
	return &SaaSHandler{repo: repo}
}

// List handles GET /api/saas
func (h *SaaSHandler) List(c *gin.Context) {
	subs, err := h.repo.FindAll()
	if err != nil {
		utils.InternalError(c, "failed to fetch saas subscriptions")
		return
	}
	utils.OK(c, subs)
}

// Summary handles GET /api/saas/summary
func (h *SaaSHandler) Summary(c *gin.Context) {
	summary, err := h.repo.Summary()
	if err != nil {
		utils.InternalError(c, "failed to fetch saas summary")
		return
	}
	utils.OK(c, summary)
}
