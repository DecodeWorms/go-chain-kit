package services

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"go-chain-kit/idgenerator"
	"gorm.io/gorm"
	"log"
	"math/big"

	"go-chain-kit/blockchain"
	"go-chain-kit/models"
)

type WalletService struct {
	db               *gorm.DB
	blockchainClient *blockchain.Client
	idGen            idgenerator.IdGenerator
}

func NewWalletService(db *gorm.DB, client *blockchain.Client, idGen idgenerator.IdGenerator) *WalletService {
	return &WalletService{
		db:               db,
		blockchainClient: client,
		idGen:            idGen,
	}
}

func (s *WalletService) CreateWallet(userID string) (*models.Wallet, error) {
	// For simplicity, we'll use the main address
	// In production, you'd generate new addresses for each user
	address := generateNewWalletAddress()
	wallet := &models.Wallet{
		ID:      s.idGen.Generate(),
		UserID:  userID,
		Address: address,
		Balance: "0",
	}

	err := s.db.Create(wallet).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %v", err)
	}
	return wallet, nil
}

func (s *WalletService) GetBalance(address string) (string, error) {
	balance, err := s.blockchainClient.GetBalance(address)
	if err != nil {
		return "", err
	}

	// Convert Wei to Ether (divide by 10^18)
	ether := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	return ether.String(), nil
}

func (s *WalletService) UpdateWalletBalance(address string) error {
	balance, err := s.GetBalance(address)
	if err != nil {
		return err
	}

	return s.db.Model(&models.Wallet{}).Where("address = ?", address).Update("balance", balance).Error
}

func (s *WalletService) GetWalletByAddress(address string) (*models.Wallet, error) {
	var wallet models.Wallet
	err := s.db.Where("address = ?", address).First(&wallet).Error
	return &wallet, err
}

func (s *WalletService) GetUserWallets(userID string) ([]models.Wallet, error) {
	var wallets []models.Wallet
	err := s.db.Where("user_id = ?", userID).Find(&wallets).Error
	return wallets, err
}

func generateNewWalletAddress() string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	return address.String()
}
