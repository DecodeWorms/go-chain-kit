package models

import (
	"time"

	_ "gorm.io/gorm"
)

type Transaction struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	TxHash      string    `json:"tx_hash" gorm:"uniqueIndex;not null"`
	FromAddress string    `json:"from_address" gorm:"not null"`
	ToAddress   string    `json:"to_address" gorm:"not null"`
	Amount      string    `json:"amount" gorm:"not null"` // Store as string to avoid precision loss
	Status      string    `json:"status" gorm:"not null"` // pending, confirmed, failed
	BlockNumber uint64    `json:"block_number"`
	GasUsed     uint64    `json:"gas_used"`
	GasPrice    string    `json:"gas_price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Wallet struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id" gorm:"not null"`
	Address   string    `json:"address" gorm:"uniqueIndex;not null"`
	Balance   string    `json:"balance" gorm:"default:'0'"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"uniqueIndex;not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Wallets   []Wallet  `json:"wallets" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserReq struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}
