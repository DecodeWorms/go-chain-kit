package services

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	_ "strconv"

	"gorm.io/gorm"

	"go-chain-kit/blockchain"
	"go-chain-kit/models"
)

type TransactionService struct {
	db               *gorm.DB
	blockchainClient *blockchain.Client
}

func NewTransactionService(db *gorm.DB, client *blockchain.Client) *TransactionService {
	return &TransactionService{
		db:               db,
		blockchainClient: client,
	}
}

func (s *TransactionService) SendTransaction(fromAddress, toAddress string, amountEther float64) (*models.Transaction, error) {
	// Convert Ether to Wei
	amount := new(big.Float).SetFloat64(amountEther)
	wei := new(big.Float).Mul(amount, big.NewFloat(1e18))
	weiInt, _ := wei.Int(nil)

	// Send transaction on blockchain
	tx, err := s.blockchainClient.SendTransaction(toAddress, weiInt)
	if err != nil {
		return nil, fmt.Errorf("failed to send blockchain transaction: %v", err)
	}

	// Save to database
	transaction := &models.Transaction{
		TxHash:      tx.Hash().Hex(),
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		Amount:      fmt.Sprintf("%.2f", amountEther),
		Status:      "pending",
		GasPrice:    tx.GasPrice().String(),
	}

	err = s.db.Create(transaction).Error
	if err != nil {
		return nil, fmt.Errorf("failed to save transaction: %v", err)
	}

	return transaction, nil
}

func (s *TransactionService) GetTransactionStatus(txHash string) (*models.Transaction, error) {
	var transaction models.Transaction
	err := s.db.Where("tx_hash = ?", txHash).First(&transaction).Error
	if err != nil {
		return nil, err
	}

	// Check blockchain for status
	receipt, err := s.blockchainClient.GetTransactionReceipt(txHash)
	if err != nil {
		// Transaction might still be pending
		return &transaction, nil
	}

	// Update transaction status
	if receipt.Status == 1 {
		transaction.Status = "confirmed"
	} else {
		transaction.Status = "failed"
	}

	transaction.BlockNumber = receipt.BlockNumber.Uint64()
	transaction.GasUsed = receipt.GasUsed

	s.db.Save(&transaction)

	return &transaction, nil
}

func (s *TransactionService) GetUserTransactions(userID uint) ([]models.Transaction, error) {
	// Get user's wallet addresses
	var wallets []models.Wallet
	err := s.db.Where("user_id = ?", userID).Find(&wallets).Error
	if err != nil {
		return nil, err
	}

	var addresses []string
	for _, wallet := range wallets {
		addresses = append(addresses, wallet.Address)
	}

	// Get transactions for these addresses
	var transactions []models.Transaction
	err = s.db.Where("from_address IN ? OR to_address IN ?", addresses, addresses).Find(&transactions).Error

	return transactions, err
}

func receiverAddress() string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	return address.String()
}
