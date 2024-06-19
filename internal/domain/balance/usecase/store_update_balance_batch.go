package usecase

import (
	"fmt"
	"sync"

	"github.com/geovanymds/balance/internal/domain/balance/dto"
	"github.com/geovanymds/balance/internal/domain/balance/entity"
)

func (uc *BalanceUseCase) StoreUpdateBalances(input *dto.CreateOrUpdateBalanceInputDto) error {
	var wg sync.WaitGroup
	wg.Add(2)

	balances := []entity.Balance{}

	balances = append(balances, *entity.NewBalance(input.AccountIDFrom, input.BalanceAccountFrom, nil, nil))
	balances = append(balances, *entity.NewBalance(input.AccountIDTo, input.BalanceAccountTo, nil, nil))

	for _, balance := range balances {
		go func() {
			defer wg.Done()
			uc.StoreUpdate(&balance)
		}()
	}

	wg.Wait()

	return nil
}

func (uc *BalanceUseCase) StoreUpdate(balance *entity.Balance) error {

	err := uc.Repository.StoreUpdate(balance)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
