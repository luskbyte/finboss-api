# Finboss API

Personal finance management API - income, expenses and investment tracking.

Go + net/http + database/sql + pgx + PostgresSQL

## Quick start

```bash
cp .env.example .env #fill in DB_PASSWORD
docker-compose up --build -d
```

API available at http://localhost:8080

### Development (without Docker)

```bash
docker-compose up postgres -d
DB_PASSWORD=your_password DB_SSLMODE=disable go run ./cmd
```

## API
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/dashboard` | Finalcial summary |
| GET/POST | `api/v1/incomes` | List / create incomes |
| GET/PUT/DELETE | `api/v1/incomes/{id}` | Get / update / delete |
| GET/POST | `/api/v1/expenses` | List / create expenses |
| GET/PUT/DELETE | `/api/v1/expenses/{id}` | Get / update / delete |
| GET/POST | `/api/v1/investments` | List / create investments |
| GET/PUT/DELETE | `/api/v1/investments/{id}` | Get / update / delete |

## Environment variables

| Variable | Default | Required |
|----------|---------|----------|
| `DB_PASSWORD` | - | yes |
| `DB_HOST` | localhost | no |
| `DB_USER` | finboss | no |
| `DB_NAME` | finboss | no |
| `DB_PORT` | 5433 | no |
| `DB_SSLMODE` | require | no |
| `PORT` | 8080 | no |
| `ALLOWED_ORIGINS` | localhost:5173,localhost:3000 | no |