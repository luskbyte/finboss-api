package database

import (
	"database/sql"
	"fmt"
	"log"
)

const schema = `
CREATE TABLE IF NOT EXISTS incomes (
	id SERIAL PRIMARY KEY,
	description TEXT,
	amount DOUBLE PRECISION NOT NULL,
	category TEXT NOT NULL,
	date TIMESTAMPZ NOT NULL,
	recurring BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
	deleted_at TIMESTAMPZ
);

CREATE TABLE IF NOT EXISTS expenses (
	id SERIAL PRIMARY KEY,
	description TEXT,
	amount DOUBLE PRECISION NOT NULL,
	category TEXT NOT NULL,
	date TIMESTAMPZ NOT NULL,
	recurring BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
	deleted_at TIMESTAMPZ
);

CREATE TABLE IF NOT EXISTS investments (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	ticker TEXT DEFAULT '',
	type TEXT NOT NULL,
	quantity DOUBLE PRECISION NOT NULL,
	buy_price DOUBLE PRECISION NOT NULL,
	current_price DOUBLE PRECISION NOT NULL DEFAULT 0,
	buy_date TIMESTAMPZ NOT NULL,
	created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
	deleted_at TIMESTAMPZ
);
`

func Migrate(db *sql.DB) {
	if _, err := db.Exec(schema); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	fmt.Println("database migrated")
}
