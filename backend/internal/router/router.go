package router

import (
	"github.com/gin-gonic/gin"

	"vendorsage/backend/internal/handlers"
)

type Handlers struct {
	Email    *handlers.EmailHandler
	Spending *handlers.SpendingHandler
	SaaS     *handlers.SaaSHandler
}

func New(h Handlers) *gin.Engine {
	r := gin.Default()

	r.Use(corsMiddleware())

	api := r.Group("/api")
	{
		api.POST("/emails/upload", h.Email.Upload)

		api.GET("/spending", h.Spending.List)
		api.GET("/spending/summary", h.Spending.Summary)

		api.GET("/saas", h.SaaS.List)
		api.GET("/saas/summary", h.SaaS.Summary)
	}

	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
