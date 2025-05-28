# 🪙 Go Blockchain Wallet Service

A simple Ethereum wallet and transaction service built in Go, using PostgreSQL for data storage, GORM for ORM, and Gin for routing.

---

## 🛠️ Tech Stack

- **Go 1.24+** – Backend implementation
- **PostgreSQL** – Primary relational database
- **GORM** – ORM for database interaction
- **Gin** – Web framework

---

## 🚀 Setup Instructions

### Prerequisites

- Go installed (>= 1.24)
- PostgreSQL running locally
- Git

---

### Clone the Repository

```bash
git clone https://github.com/DecodeWorms/go-chain-kit

Configure Environment
Create a .env file (if using environment variables):

RUN LOCALLY
go run main.go

📅 API Endpoints
👤 User
Create user

---POST api/vi/user
Request Body
{
  "user_name": "John Doe",
  "email": "john@example.com"
}
Response userID : UUID

API endpoinz for creating wallets
---POST /api/v1/wallets/?id=userID
Response{ 
   wallet_address:"0x........"
   balane : 10.00
}

---API endpoints for getting wallet balance
GET /api/v1/wallets/balance
Response{
   address : 0x......
   balance : 10.00 
}

-- API endpoint for updating balance
PUT /api/v1/wallets/balance/?address=0x......
Response{ 
   successfully updated the balace
}

💸 Transactions
Send ETH transaction
---POST /api/v1/transactions
Request body{ 
 from_address : 0x....
 to_address := 0x1....
 amount: 100.00
}

--- Check transaction status
GET /api/v1/transactions/?txHash=xxxxx
Response{ 
	TxHash : xxxxxx
	FromAddress : 0x......
	ToAddress  : 0x1.....
	Amount  : 245.0000
	Status  : successful
	GasUsed  :23455667.0000
	GasPrice  : 5755949.0000
}

--- Get Transaction by user
GET /api/v1/transactions/user/?user_id=uuid

Response{ 
{ 
 TxHash : xxxxxx
	FromAddress : 0x......
	ToAddress  : 0x1.....
	Amount  : 245.0000
	Status  : successful
	GasUsed  :23455667.0000
	GasPrice  : 5755949.0000
},
{ 
	TxHash : xxxxxx
	FromAddress : 0x......
	ToAddress  : 0x1.....
	Amount  : 245.0000
	Status  : successful
	GasUsed  :23455667.0000
	GasPrice  : 5755949.0000
}
}






