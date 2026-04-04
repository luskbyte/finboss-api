package repositories

import (
	"database/sql"
	"finboss/internal/models"
	"fmt"
)

type IncomeRepository struct {
	db *sql.DB
}

func NewIncomeRepository(db *sql.DB) *IncomeRepository {
	return &IncomeRepository{db: db}
}

func (r *IncomeRepository) FindAll() ([]models.Income, error) {
	rows, err := r.db.Query(`
		SELECT id, description, amount, category, date, recurring, created_at, updated_at 
		FROM incomes WHERE deleted_at IS NULL
		ORDER BY date DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("income.FindAll: %s", err)
	}
	defer rows.Close()

	var incomes []models.Income
	for rows.Next() {
		var income models.Income
		if err := rows.Scan(&income.ID, &income.Description, &income.Amount, &income.Category, &income.Date, &income.Recurring, &income.CreatedAt, &income.UpdatedAt); err != nil {
			return nil, fmt.Errorf("income.FindAll: %s", err)
		}
		incomes = append(incomes, income)
	}
	return incomes, nil
}

func (r *IncomeRepository) FindByID(id uint) (*models.Income, error) {
	var income models.Income
	err := r.db.QueryRow(`
		SELECT id, description, amount, category, date, recurring, created_at, updated_at 
		FROM incomes WHERE id = $1 AND deleted_at IS NULL
	`, id).Scan(&income.ID, &income.Description, &income.Amount, &income.Category, &income.Date, &income.Recurring, &income.CreatedAt, &income.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("income.FindByID: %s", err)
	}
	return &income, nil
}

func (r *IncomeRepository) Create(income *models.Income) error {
	err := r.db.QueryRow(`
		INSERT INTO incomes (description, amount, category, date, recurring, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`, income.Description, income.Amount, income.Category, income.Date, income.Recurring).Scan(&income.ID, &income.CreatedAt, &income.UpdatedAt)
	if err != nil {
		return fmt.Errorf("income.Create: %s", err)
	}
	return nil
}

func (r *IncomeRepository) Update(income *models.Income) error {
	result, err := r.db.Exec(`
		UPDATE incomes SET description = $1, amount = $2, category = $3, date = $4, recurring = $5, updated_at = NOW()
		WHERE id = $6 AND deleted_at IS NULL
	`, income.Description, income.Amount, income.Category, income.Date, income.Recurring, income.ID)
	if err != nil {
		return fmt.Errorf("income.Update: %s", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *IncomeRepository) Delete(id uint) error {
	result, err := r.db.Exec(`
		UPDATE incomes SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL
	`, id)
	if err != nil {
		return fmt.Errorf("income.Delete: %s", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *IncomeRepository) SumByMonth(year int) ([]models.MonthlyBalance, error) {
	rows, err := r.db.Query(`
		SELECT EXTRACT(MONTH FROM date) AS month, SUM(amount) 
		FROM incomes
		WHERE EXTRACT(YEAR FROM date) = $1 AND deleted_at IS NULL
		GROUP BY month
		ORDER BY month
	`, year)
	if err != nil {
		return nil, fmt.Errorf("income.SumByMonth: %s", err)
	}
	defer rows.Close()

	var monthly []models.MonthlyBalance
	for rows.Next() {
		var m models.MonthlyBalance
		var total float64
		if err := rows.Scan(&m.Month, &total); err != nil {
			return nil, fmt.Errorf("income.SumByMonth: %s", err)
		}
		m.Income = total
		monthly = append(monthly, m)
	}
	return monthly, nil
}

func (r *IncomeRepository) SumByCategory() (map[string]float64, error) {
	rows, err := r.db.Query(`
		SELECT category, SUM(amount) 
		FROM incomes
		WHERE deleted_at IS NULL
		GROUP BY category
	`)
	if err != nil {
		return nil, fmt.Errorf("income.SumByCategory: %s", err)
	}
	defer rows.Close()

	m := make(map[string]float64)
	for rows.Next() {
		var category string
		var total float64
		if err := rows.Scan(&category, &total); err != nil {
			return nil, fmt.Errorf("income.SumByCategory: %s", err)
		}
		m[category] = total
	}
	return m, nil
}
