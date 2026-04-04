package models

import (
	"time"

	"gorm.io/gorm"
)

type Income struct {
	gorm.Model
	Description string    `json:"description" gorm:"not null"`
	Amount      float64   `json:"amount" gorm:"not null"`
	Category    string    `json:"category" gorm:"not null"` // salary, freelance, bonus, other
	Date        time.Time `json: "date" gorm:"not null"`
	Recurring   bool      `json:"recurring" gorm:"default:false"`
}

type Expense struct {
	gorm.Model
	Description string    `json:"description" gorm:"not null"`
	Amount      float64   `json:"amount" gorm:"not null"`
	Category    string    `json:"category" gorm:"not null"` // rent, groceries, utilities, entertainment, other
	Date        time.Time `json: "date" gorm:"not null"`
	Recurring   bool      `json:"recurring" gorm:"default:false"`
}

type Investment struct {
	gorm.Model
	Name         string    `json:"name" gorm:"not null"`
	ticker       string    `json:"ticker"`
	Type         string    `json:"type" gorm:"not null"` // fixed, variable, investment
	Quantity     float64   `json:"quantity" gorm:"not null"`
	BuyPrice     float64   `json:"buy_price" gorm:"not null"`
	CurrentPrice float64   `json:"current_price"`
	BuyDate      time.Time `json:"buy_date" gorm:"not null"`
}

type DashboardSummary struct {
	TotalIncome       float64            `json:"total_income"`
	TotalExpense      float64            `json:"total_expense"`
	NetSavings        float64            `json:"net_savings"`
	TotalInvested     float64            `json:"total_invested"`
	PortfolioValue    float64            `json:"portfolio_value"`
	PortfolioReturn   float64            `json:"portfolio_return"`
	IncomeByCategory  map[string]float64 `json:"income_by_category"`
	ExpenseByCategory map[string]float64 `json:"expense_by_category"`
	MonthlyBalance    []MonthlyBalance   `json:"month_balance"`
}

type MonthlyBalance struct {
	Month   string  `json:"month"`
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
	Balance float64 `json:"balance"`
}
