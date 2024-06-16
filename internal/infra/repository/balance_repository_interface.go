package repository

import "github.com/geovanymds/balance/internal/domain/balance/entity"

type BalanceRepositoryInterface interface {
	StoreUpdate(balance *entity.Balance) (string, error)
	Find(id string) (*entity.Balance, error)
}
