package models

import (
	"time"
)

type ExpenseCategory string

const (
	ExpenseCategoryHousing       ExpenseCategory = "housing"
	ExpenseCategoryFood          ExpenseCategory = "food"
	ExpenseCategoryTransport     ExpenseCategory = "transport"
	ExpenseCategoryHealth        ExpenseCategory = "health"
	ExpenseCategoryEducation     ExpenseCategory = "education"
	ExpenseCategoryEntertainment ExpenseCategory = "entertainment"
	ExpenseCategoryOther         ExpenseCategory = "other"
)

func (c ExpenseCategory) IsValid() bool {
	switch c {
	case ExpenseCategoryHousing, ExpenseCategoryFood, ExpenseCategoryTransport, ExpenseCategoryHealth,
		ExpenseCategoryEducation, ExpenseCategoryEntertainment, ExpenseCategoryOther:
		return true
	}
	return false
}

type Expense struct {
	ID          uint            `json:"id"`
	Description string          `json:"description"`
	Amount      float64         `json:"amount"`
	Category    ExpenseCategory `json:"category"`
	Date        time.Time       `json: "date"`
	Recurring   bool            `json:"recurring"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}
