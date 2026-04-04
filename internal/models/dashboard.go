package models

type DashboardSummary struct {
	TotalIncome       float64            `json:"total_income"`
	TotalExpenses     float64            `json:"total_expenses"`
	Balance           float64            `json:"balance"`
	TotalInvested     float64            `json:"total_invested"`
	PortfolioValue    float64            `json:"portfolio_value"`
	PortfolioReturn   float64            `json:"portfolio_return"`
	IncomeByCategory  map[string]float64 `json:"income_by_category"`
	ExpenseByCategory map[string]float64 `json:"expense_by_category"`
	MonthlyBalance    []MonthlyBalance   `json:"monthly_balance"`
}

type MonthlyBalance struct {
	Month   string  `json:"month"`
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
	Balance float64 `json:"balance"`
}
