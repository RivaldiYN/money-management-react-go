# Money Management Application

A simple money management application built with Golang (backend), React (frontend), and PostgreSQL (database).

## Prerequisites

1. Go (1.13 or later)
2. Node.js and npm
3. PostgreSQL

## Database Setup

1. Create a PostgreSQL database:

```sql
CREATE DATABASE money_management;
```

2. The application will automatically create the necessary tables when it starts.

## Backend Setup

1. Navigate to the backend directory:

```bash
cd money-management-app/backend
```

2. Download dependencies:

```bash
go mod download
```

3. Update the database configuration in `config/config.go` if needed:

```go
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "123"
    dbname   = "money_management"
)
```

4. Run the backend:

```bash
go run main.go
```

The backend server will start at `http://localhost:8080`.

## Frontend Setup

1. Navigate to the frontend directory:

```bash
cd money-management-app/frontend
```

2. Install dependencies:

```bash
npm install
```

3. Start the frontend development server:

```bash
npm start
```

The frontend will start at `http://localhost:3000`.

## Features

- Track income and expenses
- View transaction history
- Edit and delete transactions
- View financial summary (total income, total expenses, balance)

## API Endpoints

### Transactions

- `GET /api/v1/transactions` - Get all transactions
- `GET /api/v1/transactions/:id` - Get a specific transaction
- `POST /api/v1/transactions` - Create a new transaction
- `PUT /api/v1/transactions/:id` - Update a transaction
- `DELETE /api/v1/transactions/:id` - Delete a transaction

### Summary

- `GET /api/v1/summary` - Get financial summary