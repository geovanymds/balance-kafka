package usecase

import (
	"sync"

	"github.com/geovanymds/balance/internal/domain/balance/dto"
	"github.com/geovanymds/balance/internal/domain/balance/entity"
)

func (uc *BalanceUseCase) StoreUpdateBalances(input *dto.CreateOrUpdateBalanceInputDto) error {
	var wg sync.WaitGroup
	wg.Add(2)

	balances := []entity.Balance{}

	balances = append(balances, *entity.NewBalance(input.AccountIDFrom, input.BalanceAccountFrom))
	balances = append(balances, *entity.NewBalance(input.AccountIDTo, input.BalanceAccountTo))

	errs := make(chan error, 2)

	for _, balance := range balances {
		go func() {
			errs <- uc.StoreUpdate(&balance, &wg)
		}()
	}

	for err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc *BalanceUseCase) StoreUpdate(balance *entity.Balance, wg *sync.WaitGroup) error {
	defer wg.Done()

	err := uc.Repository.StoreUpdate(balance)

	if err != nil {
		return err
	}

	return nil
}
