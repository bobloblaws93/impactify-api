package currency

import (
	"database/sql"

	"github.com/impactify-api/internal/src/models"
)

type Repository struct {
	dbClient *sql.DB
}

func NewRepository(client *sql.DB) *Repository {
	return &Repository{
		dbClient: client,
	}
}

func (r *Repository) GetCurrency(symbol string) *models.Currency {
	// sql query to get a publisher by id
	// SELECT * FROM publisher WHERE id = ?
	var currency models.Currency
	err := r.dbClient.QueryRow("SELECT id, symbol, rate FROM currency WHERE symbol = ?", symbol).Scan(&currency.ID, &currency.Symbol, &currency.Rate)
	if err != nil {
		return &models.Currency{}
	}
	return &currency
}
