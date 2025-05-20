package controllers

import (
	"net/http"
	"strconv"

	"backend/models"
	"backend/repositories"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

// TransactionController mengatur transaksi
type TransactionController struct {
	Repository *repositories.TransactionRepository
}

// NewTransactionController membuat controller transaksi baru
func NewTransactionController(repository *repositories.TransactionRepository) *TransactionController {
	return &TransactionController{Repository: repository}
}

// CreateTransaction mengatur pembuatan transaksi baru
func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	var input models.TransactionInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input data", err))
		return
	}

	transaction, err := c.Repository.CreateTransaction(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create transaction", err))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SuccessResponse("Transaction created successfully", transaction))
}

// GetAllTransactions mengatur pengambilan semua transaksi
func (c *TransactionController) GetAllTransactions(ctx *gin.Context) {
	transactions, err := c.Repository.GetAllTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve transactions", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Transactions retrieved successfully", transactions))
}

// GetTransactionByID mengatur pengambilan transaksi berdasarkan ID
func (c *TransactionController) GetTransactionByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid transaction ID", err))
		return
	}

	transaction, err := c.Repository.GetTransactionByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Transaction not found", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Transaction retrieved successfully", transaction))
}

// UpdateTransaction mengatur pembaruan transaksi
func (c *TransactionController) UpdateTransaction(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid transaction ID", err))
		return
	}

	var input models.TransactionInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input data", err))
		return
	}

	transaction, err := c.Repository.UpdateTransaction(id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update transaction", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Transaction updated successfully", transaction))
}

// DeleteTransaction mengatur penghapusan transaksi
func (c *TransactionController) DeleteTransaction(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid transaction ID", err))
		return
	}

	err = c.Repository.DeleteTransaction(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete transaction", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Transaction deleted successfully", nil))
}

// GetSummary mengatur pengambilan ringkasan transaksi
func (c *TransactionController) GetSummary(ctx *gin.Context) {
	summary, err := c.Repository.GetSummary()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve summary", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Summary retrieved successfully", summary))
}
