package models

import (
	"time"
)

type IncomeCategory string

const (
	IncomeCategorySalary    IncomeCategory = "salary"
	IncomeCategoryFreelance IncomeCategory = "freelance"
	IncomeCategoryBonus     IncomeCategory = "bonus"
	IncomeCategoryOther     IncomeCategory = "other"
)

func (c IncomeCategory) IsValid() bool {
	switch c {
	case IncomeCategorySalary, IncomeCategoryFreelance, IncomeCategoryBonus, IncomeCategoryOther:
		return true
	}
	return false
}

type Income struct {
	ID          uint           `json:"id"`
	Description string         `json:"description"`
	Amount      float64        `json:"amount"`
	Category    IncomeCategory `json:"category"`
	Date        time.Time      `json: "date"`
	Recurring   bool           `json:"recurring"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
