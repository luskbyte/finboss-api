package repositories

import (
	"database/sql"
	"finboss/internal/models"
	"fmt"
)

type ExpenseRepository struct {
	db *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (r *ExpenseRepository) FindAll() ([]models.Expense, error) {
	rows, err := r.db.Query(`
		SELECT id, description, amount, category, date, recurring, created_at, updated_at 
		FROM expenses WHERE deleted_at IS NULL
		ORDER BY date DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("expense.FindAll: %s", err)
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var expense models.Expense
		if err := rows.Scan(&expense.ID, &expense.Description, &expense.Amount, &expense.Category, &expense.Date, &expense.Recurring, &expense.CreatedAt, &expense.UpdatedAt); err != nil {
			return nil, fmt.Errorf("expense.FindAll: %s", err)
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}

func (r *ExpenseRepository) FindByID(id uint) (*models.Expense, error) {
	var expense models.Expense
	err := r.db.QueryRow(`
		SELECT id, description, amount, category, date, recurring, created_at, updated_at 
		FROM expenses WHERE id = $1 AND deleted_at IS NULL
	`, id).Scan(&expense.ID, &expense.Description, &expense.Amount, &expense.Category, &expense.Date, &expense.Recurring, &expense.CreatedAt, &expense.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("expense.FindByID: %s", err)
	}
	return &expense, nil
}

func (r *ExpenseRepository) Create(expense *models.Expense) error {
	err := r.db.QueryRow(`
		INSERT INTO expenses (description, amount, category, date, recurring, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`, expense.Description, expense.Amount, expense.Category, expense.Date, expense.Recurring).Scan(&expense.ID, &expense.CreatedAt, &expense.UpdatedAt)
	if err != nil {
		return fmt.Errorf("expense.Create: %s", err)
	}
	return nil
}

func (r *ExpenseRepository) Update(expense *models.Expense) error {
	result, err := r.db.Exec(`
		UPDATE expenses
		SET description = $1, amount = $2, category = $3, date = $4, recurring = $5, updated_at = NOW()
		WHERE id = $6 AND deleted_at IS NULL
	`, expense.Description, expense.Amount, expense.Category, expense.Date, expense.Recurring, expense.ID)
	if err != nil {
		return fmt.Errorf("expense.Update: %s", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *ExpenseRepository) Delete(id uint) error {
	result, err := r.db.Exec(`
		UPDATE expenses SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL
	`, id)
	if err != nil {
		return fmt.Errorf("expense.Delete: %s", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *ExpenseRepository) SumByMonth(year int) ([]models.MonthlyBalance, error) {
	rows, err := r.db.Query(`
		SELECT TO_CHAR(date, 'YYYY-MM') as month, SUM(amount) as total
		FROM expenses
		WHERE EXTRACT(YEAR FROM date) = $1 AND deleted_at IS NULL
		GROUP BY month ORDER BY month
	`, year)
	if err != nil {
		return nil, fmt.Errorf("expense.SumByMonth: %s", err)
	}
	defer rows.Close()

	var monthly []models.MonthlyBalance
	for rows.Next() {
		var m models.MonthlyBalance
		var total float64
		if err := rows.Scan(&m.Month, &total); err != nil {
			return nil, fmt.Errorf("expense.SumByMonth: %s", err)
		}
		m.Expense = total
		monthly = append(monthly, m)
	}
	return monthly, nil
}

func (r *ExpenseRepository) SumByCategory() (map[string]float64, error) {
	rows, err := r.db.Query(`
		SELECT category, SUM(amount) as total
		FROM expenses WHERE deleted_at IS NULL
		GROUP BY category
	`)
	if err != nil {
		return nil, fmt.Errorf("expense.SumByCategory: %s", err)
	}
	defer rows.Close()

	m := make(map[string]float64)
	for rows.Next() {
		var category string
		var total float64
		if err := rows.Scan(&category, &total); err != nil {
			return nil, fmt.Errorf("expense.SumByCategory: %s", err)
		}
		m[category] = total
	}
	return m, nil
}
