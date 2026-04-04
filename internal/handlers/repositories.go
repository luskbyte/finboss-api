package handlers

import (
	"finboss/internal/models"
)

type IncomeRepository interface {
	FindAll() ([]models.Income, error)
	FindByID(id uint) (*models.Income, error)
	Create(income *models.Income) error
	Update(income *models.Income) error
	Delete(id uint) error
	SumByMonth(year int) ([]models.MonthlyBalance, error)
	SumByCategory() (map[string]float64, error)
}

type ExpenseRepository interface {
	FindAll() ([]models.Expense, error)
	FindByID(id uint) (*models.Expense, error)
	Create(expense *models.Expense) error
	Update(expense *models.Expense) error
	Delete(id uint) error
	SumByMonth(year int) ([]models.MonthlyBalance, error)
	SumByCategory() (map[string]float64, error)
}

type InvestmentRepository interface {
	FindAll() ([]models.Investment, error)
	FindByID(id uint) (*models.Investment, error)
	Create(investment *models.Investment) error
	Update(investment *models.Investment) error
	Delete(id uint) error
	GetPortfolioSummary() (totalInvested, portfolioValue float64, err error)
}
