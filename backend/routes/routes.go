package routes

import (
	"database/sql"
	"time"

	"backend/controllers"
	"backend/repositories"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter mengatur semua rute dan middleware aplikasi.
func SetupRouter(db *sql.DB) *gin.Engine {
	// Inisialisasi Gin
	router := gin.Default()

	// Konfigurasi CORS agar frontend (localhost:3000) bisa akses backend
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Inisialisasi repository dan controller
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionController := controllers.NewTransactionController(transactionRepo)

	// Buat grup API v1
	api := router.Group("/api/v1")
	{
		// Rute transaksi
		transactions := api.Group("/transactions")
		{
			transactions.POST("", transactionController.CreateTransaction)
			transactions.GET("", transactionController.GetAllTransactions)
			transactions.GET("/:id", transactionController.GetTransactionByID)
			transactions.PUT("/:id", transactionController.UpdateTransaction)
			transactions.DELETE("/:id", transactionController.DeleteTransaction)
		}

		// Rute ringkasan
		api.GET("/summary", transactionController.GetSummary)
	}

	return router
}
