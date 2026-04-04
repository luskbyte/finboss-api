package models

import (
	"time"

	"gorm.io/gorm"
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
	gorm.Model
	Description string          `json:"description" gorm:"not null"`
	Amount      float64         `json:"amount" gorm:"not null"`
	Category    ExpenseCategory `json:"category" gorm:"not null"`
	Date        time.Time       `json: "date" gorm:"not null"`
	Recurring   bool            `json:"recurring" gorm:"default:false"`
}
