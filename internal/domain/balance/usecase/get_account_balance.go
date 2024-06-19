package usecase

import (
	"github.com/geovanymds/balance/internal/domain/balance/dto"
	"github.com/geovanymds/balance/internal/domain/balance/entity"
)

func (uc *BalanceUseCase) GetBalanceByAccountId(input *dto.GetAccountBalanceDto) (*entity.Balance, error) {
	balance, err := uc.Repository.Find(input.AccountId)

	if err != nil {
		return nil, err
	}

	return entity.NewBalance(balance.AccountID, balance.Value, balance.CreatedAt, balance.UpdatedAt), nil

}
