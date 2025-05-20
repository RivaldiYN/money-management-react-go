package repositories

import (
	"database/sql"
	"log"

	"backend/models"
)

// TransactionRepository untuk handle database transaksi
type TransactionRepository struct {
	DB *sql.DB
}

// NewTransactionRepository buat baru repository transaksi
func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

// CreateTransaction untuk tambah baru transaksi ke database
func (r *TransactionRepository) CreateTransaction(input models.TransactionInput) (models.Transaction, error) {
	var transaction models.Transaction

	err := r.DB.QueryRow(`
		INSERT INTO transactions (type, amount, category, description)
		VALUES ($1, $2, $3, $4)
		RETURNING id, type, amount, category, description, date
	`, input.Type, input.Amount, input.Category, input.Description).Scan(
		&transaction.ID,
		&transaction.Type,
		&transaction.Amount,
		&transaction.Category,
		&transaction.Description,
		&transaction.Date,
	)

	return transaction, err
}

// GetAllTransactions mengambil semua transaksi dari database
func (r *TransactionRepository) GetAllTransactions() ([]models.Transaction, error) {
	rows, err := r.DB.Query(`
		SELECT id, type, amount, category, description, date
		FROM transactions
		ORDER BY date DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		if err := rows.Scan(
			&transaction.ID,
			&transaction.Type,
			&transaction.Amount,
			&transaction.Category,
			&transaction.Description,
			&transaction.Date,
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// GetTransactionByID mengambil transaksi berdasarkan id
func (r *TransactionRepository) GetTransactionByID(id int) (models.Transaction, error) {
	var transaction models.Transaction

	err := r.DB.QueryRow(`
		SELECT id, type, amount, category, description, date
		FROM transactions
		WHERE id = $1
	`, id).Scan(
		&transaction.ID,
		&transaction.Type,
		&transaction.Amount,
		&transaction.Category,
		&transaction.Description,
		&transaction.Date,
	)

	return transaction, err
}

// UpdateTransaction mengumbah transaksi yang sudah ada
func (r *TransactionRepository) UpdateTransaction(id int, input models.TransactionInput) (models.Transaction, error) {
	var transaction models.Transaction

	err := r.DB.QueryRow(`
		UPDATE transactions
		SET type = $1, amount = $2, category = $3, description = $4
		WHERE id = $5
		RETURNING id, type, amount, category, description, date
	`, input.Type, input.Amount, input.Category, input.Description, id).Scan(
		&transaction.ID,
		&transaction.Type,
		&transaction.Amount,
		&transaction.Category,
		&transaction.Description,
		&transaction.Date,
	)

	return transaction, err
}

// DeleteTransaction menghapus transaksi berdasarkan id
func (r *TransactionRepository) DeleteTransaction(id int) error {
	_, err := r.DB.Exec("DELETE FROM transactions WHERE id = $1", id)
	return err
}

// GetSummary mengambil ringkasan transaksi dari database
func (r *TransactionRepository) GetSummary() (models.Summary, error) {
	var summary models.Summary

	// Mendapatkan total pendapatan
	err := r.DB.QueryRow(`
		SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE type = 'income'
	`).Scan(&summary.TotalIncome)
	if err != nil {
		log.Printf("Error getting total income: %v", err)
		return summary, err
	}

	// Mendapatkan total pengeluaran
	err = r.DB.QueryRow(`
		SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE type = 'expense'
	`).Scan(&summary.TotalExpense)
	if err != nil {
		log.Printf("Error getting total expense: %v", err)
		return summary, err
	}

	// Menghitung saldo
	summary.Balance = summary.TotalIncome - summary.TotalExpense

	return summary, nil
}
