package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	address    common.Address
	networkID  *big.Int
}

func NewClient(rpcURL string, privateKeyHex string, networkID int64) (*Client, error) {
	// Connect to Ethereum client
	//Connecting to an Ethereum-based blockchain
	log.Println("Connecting to an Ethereum-based blockchain")
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	// Load private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	// Connected to an Ethereum-based blockchain client
	log.Println("Connected to an Ethereum-based blockchain")
	return &Client{
		client:     client,
		privateKey: privateKey,
		publicKey:  publicKeyECDSA,
		address:    address,
		networkID:  big.NewInt(networkID),
	}, nil
}

func (c *Client) GetBalance(address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := c.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (c *Client) GetLatestBlockNumber() (uint64, error) {
	header, err := c.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	return header.Number.Uint64(), nil
}

func (c *Client) SendTransaction(toAddress string, amount *big.Int) (*types.Transaction, error) {
	nonce, err := c.client.PendingNonceAt(context.Background(), c.address)
	if err != nil {
		return nil, err
	}

	gasLimit := uint64(21000) // Standard ETH transfer
	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	toAddr := common.HexToAddress(toAddress)
	tx := types.NewTransaction(nonce, toAddr, amount, gasLimit, gasPrice, nil)
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(c.networkID), c.privateKey)
	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(c.networkID), c.privateKey)
	if err != nil {
		return nil, err
	}

	err = c.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (c *Client) GetTransactionReceipt(txHash string) (*types.Receipt, error) {
	hash := common.HexToHash(txHash)
	receipt, err := c.client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (c *Client) GetAddress() string {
	return c.address.Hex()
}

func (c *Client) Close() {
	c.client.Close()
}
