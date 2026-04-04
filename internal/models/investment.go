package models

import (
	"time"
)

type InvestmentType string

const (
	InvestmentTypeStock       InvestmentType = "stock"
	InvestmentTypeFixedIncome InvestmentType = "fixed_income"
	InvestmentTypeFund        InvestmentType = "fund"
	InvestmentTypeCrypto      InvestmentType = "crypto"
	InvestmentTypeOther       InvestmentType = "other"
)

func (t InvestmentType) IsValid() bool {
	switch t {
	case InvestmentTypeStock, InvestmentTypeFixedIncome, InvestmentTypeFund, InvestmentTypeCrypto, InvestmentTypeOther:
		return true
	}
	return false
}

type Investment struct {
	ID           uint           `json:"id"`
	Name         string         `json:"name" gorm:"not null"`
	Ticker       string         `json:"ticker"`
	Type         InvestmentType `json:"type" gorm:"not null"`
	Quantity     float64        `json:"quantity" gorm:"not null"`
	BuyPrice     float64        `json:"buy_price" gorm:"not null"`
	CurrentPrice float64        `json:"current_price"`
	BuyDate      time.Time      `json:"buy_date" gorm:"not null"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}
