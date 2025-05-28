package main

import (
	"go-chain-kit/idgenerator"
	"log"

	"github.com/gin-gonic/gin"

	"go-chain-kit/blockchain"
	"go-chain-kit/config"
	"go-chain-kit/database"
	"go-chain-kit/handler"
	"go-chain-kit/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize blockchain client
	blockchainClient, err := blockchain.NewClient(cfg.EthereumRPCURL, cfg.EthereumPrivateKey, cfg.EthereumNetworkID)
	if err != nil {
		log.Fatal("Failed to initialize blockchain client:", err)
	}
	defer blockchainClient.Close()

	// Initialize database
	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Initialize services
	walletService := services.NewWalletService(db.DB, blockchainClient, idgenerator.New())
	transactionService := services.NewTransactionService(db.DB, blockchainClient)
	userService := services.NewUser(db.DB, idgenerator.New())

	// Initialize handlers
	walletHandler := handlers.NewWalletHandler(walletService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	userHandler := handlers.NewUserHandler(userService)

	// Setup Gin router
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API routes
	api := r.Group("/api/v1")
	{
		// Wallet routes
		api.POST("/user", userHandler.Create())
		api.POST("/wallets", walletHandler.CreateWallet)
		api.GET("/wallets", walletHandler.GetUserWallets)
		api.GET("/wallets/balance", walletHandler.GetBalance)
		api.PUT("/wallets/balance", walletHandler.UpdateBalance)

		// Transaction routes
		api.POST("/transactions", transactionHandler.SendTransaction)
		api.GET("/transactions", transactionHandler.GetTransactionStatus)
		api.GET("/users/:userId/transactions", transactionHandler.GetUserTransactions)
	}

	log.Printf("Server starting on %s:%s", cfg.APIHost, cfg.APIPort)
	log.Fatal(r.Run(":" + cfg.APIPort))
}
