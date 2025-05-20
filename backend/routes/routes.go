package routes

import (
	"database/sql"
	"time"

	"backend/controllers"
	"backend/repositories"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter mengatur router untuk aplikasi
func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	// Konfigurasi CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Memuat repository dan controller
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionController := controllers.NewTransactionController(transactionRepo)

	// API v1 group
	v1 := router.Group("/api/v1")
	{
		transactions := v1.Group("/transactions")
		{
			transactions.POST("", transactionController.CreateTransaction)
			transactions.GET("", transactionController.GetAllTransactions)
			transactions.GET("/:id", transactionController.GetTransactionByID)
			transactions.PUT("/:id", transactionController.UpdateTransaction)
			transactions.DELETE("/:id", transactionController.DeleteTransaction)
		}

		// Endpoint untuk mendapatkan ringkasan transaksi
		v1.GET("/summary", transactionController.GetSummary)
	}

	return router
}
