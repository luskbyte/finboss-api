package main

import (
	"context"
	"finboss/internal/database"
	"finboss/internal/handlers"
	"finboss/internal/middleware"
	"finboss/internal/repositories"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db := database.Connect()
	defer db.Close()

	database.Migrate(db)

	incomeRepo := repositories.NewIncomeRepository(db)
	expenseRepo := repositories.NewExpenseRepository(db)
	investmentRepo := repositories.NewInvestmentRepository(db)

	incomeHandler := handlers.NewIncomeHandler(incomeRepo)
	expenseHandler := handlers.NewExpenseHandler(expenseRepo)
	investmentHandler := handlers.NewInvestmentHandler(investmentRepo)
	dashboardHandler := handlers.NewDashboardHandler(incomeRepo, expenseRepo, investmentRepo)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/dashboard", dashboardHandler.Summary)

	mux.HandleFunc("GET /api/v1/incomes", incomeHandler.List)
	mux.HandleFunc("GET /api/v1/incomes/{id}", incomeHandler.Get)
	mux.HandleFunc("POST /api/v1/incomes", incomeHandler.Create)
	mux.HandleFunc("PUT /api/v1/incomes/{id}", incomeHandler.Update)
	mux.HandleFunc("DELETE /api/v1/incomes/{id}", incomeHandler.Delete)

	mux.HandleFunc("GET /api/v1/expenses", expenseHandler.List)
	mux.HandleFunc("GET /api/v1/expenses/{id}", expenseHandler.Get)
	mux.HandleFunc("POST /api/v1/expenses", expenseHandler.Create)
	mux.HandleFunc("PUT /api/v1/expenses/{id}", expenseHandler.Update)
	mux.HandleFunc("DELETE /api/v1/expenses/{id}", expenseHandler.Delete)

	mux.HandleFunc("GET /api/v1/investments", investmentHandler.List)
	mux.HandleFunc("GET /api/v1/investments/{id}", investmentHandler.Get)
	mux.HandleFunc("POST /api/v1/investments", investmentHandler.Create)
	mux.HandleFunc("PUT /api/v1/investments/{id}", investmentHandler.Update)
	mux.HandleFunc("DELETE /api/v1/investments/{id}", investmentHandler.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	listenAndServe(port, middleware.CORS(mux))
}

func listenAndServe(port string, handler http.Handler) {
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Printf("Server running on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("server stopped")
}
