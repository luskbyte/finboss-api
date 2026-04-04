package models

import (
	"time"

	"gorm.io/gorm"
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
	gorm.Model
	Description string         `json:"description" gorm:"not null"`
	Amount      float64        `json:"amount" gorm:"not null"`
	Category    IncomeCategory `json:"category" gorm:"not null"`
	Date        time.Time      `json: "date" gorm:"not null"`
	Recurring   bool           `json:"recurring" gorm:"default:false"`
}
