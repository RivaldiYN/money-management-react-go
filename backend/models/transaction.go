package models

import (
	"time"
)

// Database Transaction
type Transaction struct {
	ID          uint      `json:"id"`
	Type        string    `json:"type"` // 'income' or 'expense'
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

// Database Transaction Input
type TransactionInput struct {
	Type        string  `json:"type" binding:"required,oneof=income expense"`
	Amount      float64 `json:"amount" binding:"required,gt=0"`
	Category    string  `json:"category" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

// Database Summary
type Summary struct {
	TotalIncome  float64 `json:"total_income"`
	TotalExpense float64 `json:"total_expense"`
	Balance      float64 `json:"balance"`
}
