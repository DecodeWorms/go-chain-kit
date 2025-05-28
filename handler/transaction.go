package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-chain-kit/services"
)

type TransactionHandler struct {
	transactionService *services.TransactionService
}

type SendTransactionRequest struct {
	FromAddress string  `json:"from_address" binding:"required"`
	ToAddress   string  `json:"to_address" binding:"required"`
	Amount      float64 `json:"amount" binding:"required,gt=0"`
}

func NewTransactionHandler(transactionService *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) SendTransaction(c *gin.Context) {
	var req SendTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := h.transactionService.SendTransaction(req.FromAddress, req.ToAddress, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

func (h *TransactionHandler) GetTransactionStatus(c *gin.Context) {
	txHash := c.Query("txHash")
	if txHash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url parameter is missing"})
	}

	transaction, err := h.transactionService.GetTransactionStatus(txHash)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) GetUserTransactions(c *gin.Context) {
	userIDStr := c.Param("userId")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	transactions, err := h.transactionService.GetUserTransactions(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
