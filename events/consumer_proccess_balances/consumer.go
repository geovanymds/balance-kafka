package consumer_proccess_balances

import (
	"github.com/geovanymds/balance/internal/domain/balance/usecase"
)

type ConsumerBalances struct {
	uc *usecase.BalanceUseCase
}

func NewConsumerBalances(uc *usecase.BalanceUseCase) *ConsumerBalances {
	return &ConsumerBalances{uc: uc}
}
