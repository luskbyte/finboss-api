package handlers

import (
	"finboss/internal/models"
	"net/http"
	"strconv"
	"time"
)

type DashboardHandler struct {
	incomeRepo     IncomeRepository
	expenseRepo    ExpenseRepository
	InvestmentRepo InvestmentRepository
}

func NewDashboardHandler(incomeRepo IncomeRepository, expenseRepo ExpenseRepository, investmentRepo InvestmentRepository) *DashboardHandler {
	return &DashboardHandler{
		incomeRepo:     incomeRepo,
		expenseRepo:    expenseRepo,
		InvestmentRepo: investmentRepo,
	}
}

func (h *DashboardHandler) Summary(w http.ResponseWriter, r *http.Request) {
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		year = time.Now().Year()
	}

	incomes, err := h.incomeRepo.FindAll()
	if err != nil {
		respondDBError(w, err)
		return
	}
	expenses, err := h.expenseRepo.FindAll()
	if err != nil {
		respondDBError(w, err)
		return
	}

	var totalIncome, totalExpenses float64
	for _, i := range incomes {
		totalIncome += i.Amount
	}
	for _, e := range expenses {
		totalExpenses += e.Amount
	}

	incomeByCategory, err := h.incomeRepo.SumByCategory()
	if err != nil {
		respondDBError(w, err)
		return
	}
	expenseByCategory, err := h.expenseRepo.SumByCategory()
	if err != nil {
		respondDBError(w, err)
		return
	}

	totalInvested, portfolioValue, err := h.InvestmentRepo.GetPortfolioSummary()
	if err != nil {
		respondDBError(w, err)
		return
	}
	portfolioReturn := 0.0
	if totalInvested > 0 {
		portfolioReturn = (totalInvested - totalIncome) / totalInvested * 100
	}

	monthlyIncome, err := h.incomeRepo.SumByMonth(year)
	if err != nil {
		respondDBError(w, err)
		return
	}
	monthlyExpense, err := h.expenseRepo.SumByMonth(year)
	if err != nil {
		respondDBError(w, err)
		return
	}

	monthlyMap := make(map[string]*models.MonthlyBalance)
	for _, m := range monthlyIncome {
		if _, ok := monthlyMap[m.Month]; !ok {
			monthlyMap[m.Month] = &models.MonthlyBalance{Month: m.Month}
		}
		monthlyMap[m.Month].Income = m.Income
	}

	for _, m := range monthlyExpense {
		if _, ok := monthlyMap[m.Month]; !ok {
			monthlyMap[m.Month] = &models.MonthlyBalance{Month: m.Month}
		}
		monthlyMap[m.Month].Expense = m.Expense
	}

	var monthlyBalance []models.MonthlyBalance
	for _, m := range monthlyMap {
		m.Balance = m.Income - m.Expense
		monthlyBalance = append(monthlyBalance, *m)
	}

	summary := models.DashboardSummary{
		TotalIncome:       totalIncome,
		TotalExpenses:     totalExpenses,
		Balance:           totalIncome - totalExpenses,
		TotalInvested:     totalInvested,
		PortfolioValue:    portfolioValue,
		PortfolioReturn:   portfolioReturn,
		IncomeByCategory:  incomeByCategory,
		ExpenseByCategory: expenseByCategory,
		MonthlyBalance:    monthlyBalance,
	}

	writeJSON(w, http.StatusOK, summary)
}
