package repositories

import (
	"database/sql"
	"finboss/internal/models"
	"fmt"
)

type InvestmentRepository struct {
	db *sql.DB
}

func NewInvestmentRepository(db *sql.DB) *InvestmentRepository {
	return &InvestmentRepository{db: db}
}

func (r *InvestmentRepository) FindAll() ([]models.Investment, error) {
	rows, err := r.db.Query(`
		SELECT id, name, ticker, type, quantity, buy_price, current_price, buy_date, created_at, updated_at 
		FROM investments WHERE deleted_at IS NULL
		ORDER BY date DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("investment.FindAll: %s", err)
	}
	defer rows.Close()

	var investments []models.Investment
	for rows.Next() {
		var investment models.Investment
		if err := rows.Scan(&investment.ID, &investment.Name, &investment.Ticker, &investment.Type, &investment.Quantity, &investment.BuyPrice, &investment.CurrentPrice, &investment.BuyDate, &investment.CreatedAt, &investment.UpdatedAt); err != nil {
			return nil, fmt.Errorf("investment.FindAll: %s", err)
		}
		investments = append(investments, investment)
	}
	return investments, nil
}

func (r *InvestmentRepository) FindByID(id uint) (*models.Investment, error) {
	var investment models.Investment
	err := r.db.QueryRow(`
		SELECT id, name, ticker, type, quantity, buy_price, current_price, buy_date, created_at, updated_at 
		FROM investments WHERE id = $1 AND deleted_at IS NULL
	`, id).Scan(&investment.ID, &investment.Name, &investment.Ticker, &investment.Type, &investment.Quantity, &investment.BuyPrice, &investment.CurrentPrice, &investment.BuyDate, &investment.CreatedAt, &investment.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("investment.FindByID: %s", err)
	}
	return &investment, nil
}

func (r *InvestmentRepository) Create(investment *models.Investment) error {
	err := r.db.QueryRow(`
		INSERT INTO investments (name, ticker, type, quantity, buy_price, current_price, buy_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`, investment.Name, investment.Ticker, investment.Type, investment.Quantity, investment.BuyPrice, investment.CurrentPrice).Scan(&investment.ID, &investment.CreatedAt, &investment.UpdatedAt)
	if err != nil {
		return fmt.Errorf("investment.Create: %s", err)
	}
	return nil
}

func (r *InvestmentRepository) Update(investment *models.Investment) error {
	_, err := r.db.Exec(`
		UPDATE investments SET name = $1, ticker = $2, type = $3, quantity = $4, buy_price = $5, current_price = $6, buy_date = $7, updated_at = NOW()
		WHERE id = $8 AND deleted_at IS NULL
	`, investment.Name, investment.Ticker, investment.Type, investment.Quantity, investment.BuyPrice, investment.CurrentPrice, investment.BuyDate, investment.ID)
	if err != nil {
		return fmt.Errorf("investment.Update: %s", err)
	}
	return nil
}

func (r *InvestmentRepository) Delete(id uint) error {
	_, err := r.db.Exec(`
		UPDATE investments SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL
	`, id)
	if err != nil {
		return fmt.Errorf("investment.Delete: %s", err)
	}
	return nil
}

func (r *InvestmentRepository) GetPortfolioSummary() (totalInvested, portfolioValue float64, err error) {
	err = r.db.QueryRow(`
		SELECT
			COALESCE(SUM(quantity * buy_price), 0),
			COALESCE(SUM(quantity * COALESCE(NULLIF(current_price, 0), buy_price)), 0)
		FROM investments
		WHERE deleted_at IS NULL
	`).Scan(&totalInvested, &portfolioValue)
	if err != nil {
		return 0, 0, fmt.Errorf("investment.GetPortfolioSummary: %s", err)
	}
	return totalInvested, portfolioValue, nil
}
