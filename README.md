
# 🌟 Test EWallet API 🌟

Welcome to the Test EWallet API, a Go-developed application designed to manage digital wallets with transaction processing. This API allows users to create wallets, transfer funds, view transaction histories, and check balances with RESTful endpoints.

## ✨ Features

- **Wallet Creation**: Initiate with a balance of 100.0 units.
- **Fund Transfers**: Facilitate funds transfer between wallets securely.
- **Transaction History**: Access detailed records of transactions.
- **Balance Inquiry**: Review the current balance effortlessly.

## 🚀 Quick Start

1. **Clone the repository:**

```bash
git clone https://github.com/Sysleec/TestEWallet.git
```

2. **Navigate into the project directory:**

```bash
cd TestEWallet
```

3. **Run the application:**

```bash
docker-compose up -d
```

## 🌐 API Endpoints & Examples

All endpoints can be tested at `45.8.97.234:8081`.

### Create Wallet

- `POST /api/v1/wallet`
- **Request**: No body required.
- **Response**:
  ```json
  {
    "id": "wallet_id_generated_by_server",
    "balance": 100.0
  }
  ```

### Transfer Funds

- `POST /api/v1/wallet/{walletId}/send`
- **Request**:
  ```json
  {
    "to": "recipient_wallet_id",
    "amount": 50.0
  }
  ```
- **Response**: Status 200 OK if successful.

### Transaction History

- `GET /api/v1/wallet/{walletId}/history`
- **Response**:
  ```json
  [
    {
      "time": "RFC 3339 formatted date-time",
      "from": "sender_wallet_id",
      "to": "recipient_wallet_id",
      "amount": 25.0
    }
  ]
  ```

### Check Balance

- `GET /api/v1/wallet/{walletId}`
- **Response**:
  ```json
  {
    "id": "wallet_id",
    "balance": current_balance
  }
  ```

## 📂 Repository Structure

```sh
└── TestEWallet/
    ├── Dockerfile
    ├── Makefile
    ├── README.md
    ├── cmd
    │   └── main.go
    ├── docker-compose.yaml
    ├── go.mod
    ├── go.sum
    ├── go.work.sum
    ├── internal
    │   ├── api
    │   │   └── wallet
    │   │       ├── create_wallet.go
    │   │       ├── get_history.go
    │   │       ├── get_wallet.go
    │   │       ├── send_money.go
    │   │       ├── service.go
    │   │       └── tests
    │   │           ├── create_wallet_test.go
    │   │           ├── get_history_test.go
    │   │           ├── get_wallet_test.go
    │   │           └── send_money_test.go
    │   ├── converter
    │   │   └── wallet.go
    │   ├── model
    │   │   ├── errors.go
    │   │   └── wallet.go
    │   ├── repository
    │   │   ├── generate.go
    │   │   ├── mocks
    │   │   │   └── wallet_repository_minimock.go
    │   │   ├── repository.go
    │   │   └── wallet
    │   │       ├── converter
    │   │       │   └── wallet.go
    │   │       ├── create_wallet.go
    │   │       ├── get_history.go
    │   │       ├── get_wallet.go
    │   │       ├── repository.go
    │   │       ├── send_money.go
    │   │       └── sqlc
    │   │           ├── db.go
    │   │           ├── history.sql.go
    │   │           ├── models.go
    │   │           └── wallet.sql.go
    │   ├── service
    │   │   ├── generate.go
    │   │   ├── mocks
    │   │   │   └── wallet_service_minimock.go
    │   │   ├── service.go
    │   │   └── wallet
    │   │       ├── create_wallet.go
    │   │       ├── get_history.go
    │   │       ├── get_wallet.go
    │   │       ├── send_money.go
    │   │       ├── service.go
    │   │       └── tests
    │   │           ├── create_wallet_test.go
    │   │           ├── get_history_test.go
    │   │           ├── get_wallet_test.go
    │   │           └── send_money_test.go
    │   └── utils
    │       └── json.go
    ├── migration.Dockerfile
    ├── migration.sh
    ├── sql
    │   ├── fs.go
    │   ├── queries
    │   │   ├── history.sql
    │   │   └── wallet.sql
    │   └── schema
    │       ├── 001_wallets.sql
    │       └── 002_history.sql
    └── sqlc.yaml
```

## 💡 Contributing

Your contributions are welcome! Please submit a pull request or create an issue for any features, bug fixes, or improvements.

## ❌ No License

This project does not specify a license at the moment.

Thank you for checking the Test EWallet API! 
