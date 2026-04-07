package main

import (
	"log"

	"vendorsage/backend/internal/config"
	"vendorsage/backend/internal/database"
	"vendorsage/backend/internal/handlers"
	"vendorsage/backend/internal/repository"
	"vendorsage/backend/internal/router"
	"vendorsage/backend/internal/services"
)

func main() {
	cfg := config.Load()

	if cfg.AnthropicKey == "" {
		log.Fatal("ANTHROPIC_API_KEY is not set — check your .env file")
	}
	log.Printf("Anthropic key loaded: %s...%s", cfg.AnthropicKey[:8], cfg.AnthropicKey[len(cfg.AnthropicKey)-4:])

	db := database.Connect(cfg)

	// Repositories
	emailRepo := repository.NewEmailRepository(db)
	expenseRepo := repository.NewExpenseRepository(db)
	saasRepo := repository.NewSaaSRepository(db)

	// Services
	aiService := services.NewAIService(cfg.AnthropicKey)
	processor := services.NewEmailProcessor(aiService, emailRepo, expenseRepo, saasRepo)

	// Handlers
	h := router.Handlers{
		Email:    handlers.NewEmailHandler(processor),
		Spending: handlers.NewSpendingHandler(expenseRepo),
		SaaS:     handlers.NewSaaSHandler(saasRepo),
	}

	r := router.New(h)

	log.Printf("Server starting on :%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
