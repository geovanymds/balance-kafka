package repository

import (
	"github.com/geovanymds/balance/internal/domain/balance/entity"
	database "github.com/geovanymds/balance/internal/infra/db"
)

type BalanceRepository struct {
	dbClient *database.ClientDB
}

func NewBalanceRepository(db *database.ClientDB) *BalanceRepository {
	return &BalanceRepository{
		db,
	}
}

func (r *BalanceRepository) StoreUpdate(balance *entity.Balance) error {
	stmt, err := r.dbClient.DB.Prepare("INSERT INTO balances (id, account_id, value, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE value=?;")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		balance.ID,
		balance.AccountID,
		balance.Value,
		balance.CreatedAt,
		balance.UpdatedAt,
		balance.Value,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *BalanceRepository) Find(accountId string) (*entity.Balance, error) {
	balance := &entity.Balance{}
	stmt, err := r.dbClient.DB.Prepare("SELECT id, account_id, value, created_at, updated_at FROM balances WHERE account_id = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(accountId)

	if err := row.Scan(&balance.ID, &balance.AccountID, &balance.Value, &balance.CreatedAt, &balance.UpdatedAt); err != nil {
		return nil, err
	}
	return balance, nil
}
