package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"go-chain-kit/services"
)

type WalletHandler struct {
	walletService *services.WalletService
}

func NewWalletHandler(walletService *services.WalletService) *WalletHandler {
	return &WalletHandler{walletService: walletService}
}

func (h *WalletHandler) CreateWallet(c *gin.Context) {
	userID := c.Query("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameter"})
		return
	}

	wallet, err := h.walletService.CreateWallet(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, wallet)
}

func (h *WalletHandler) GetBalance(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "address parameter is missing"})
		return
	}

	balance, err := h.walletService.GetBalance(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"address": address, "balance": balance})
}

func (h *WalletHandler) UpdateBalance(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wallet address on a blockchain is missing"})
		return
	}

	err := h.walletService.UpdateWalletBalance(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Balance updated successfully"})
}

func (h *WalletHandler) GetUserWallets(c *gin.Context) {
	userID := c.Query("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing parameter"})
	}
	wallets, err := h.walletService.GetUserWallets(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wallets)
}
