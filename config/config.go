package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// Blockchain Config
	EthereumRPCURL     string
	EthereumNetworkID  int64
	EthereumPrivateKey string

	// Database Config
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// API Config
	APIPort string
	APIHost string

	// Security
	JWTSecret string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	networkID, _ := strconv.ParseInt(os.Getenv("ETHEREUM_NETWORK_ID"), 10, 64)

	return &Config{
		EthereumRPCURL:     os.Getenv("ETHEREUM_RPC_URL"),
		EthereumNetworkID:  networkID,
		EthereumPrivateKey: os.Getenv("ETHEREUM_PRIVATE_KEY"),
		DBHost:             os.Getenv("DB_HOST"),
		DBPort:             os.Getenv("DB_PORT"),
		DBUser:             os.Getenv("DB_USER"),
		DBPassword:         os.Getenv("DB_PASSWORD"),
		DBName:             os.Getenv("DB_NAME"),
		APIPort:            os.Getenv("API_PORT"),
		APIHost:            os.Getenv("API_HOST"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
	}
}
