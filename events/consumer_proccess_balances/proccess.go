package consumer_proccess_balances

import (
	"fmt"

	"github.com/geovanymds/balance/internal/domain/balance/dto"
)

type MessageBalanceUpdated struct {
	AccountIdFrom      string  `json:"account_id_from" validate:"required"`
	AccountIdTo        string  `json:"account_id_to" validate:"required"`
	BalanceAccountFrom float64 `json:"balance_account_id_from" validate:"required"`
	BalanceAccountTo   float64 `json:"balance_account_id_to" validate:"required"`
}

func (c *ConsumerBalances) Proccess(msg *MessageBalanceUpdated) error {
	createOrUpdateBalances := &dto.CreateOrUpdateBalanceInputDto{
		AccountIDFrom:      msg.AccountIdFrom,
		AccountIDTo:        msg.AccountIdTo,
		BalanceAccountFrom: msg.BalanceAccountFrom,
		BalanceAccountTo:   msg.BalanceAccountTo,
	}

	err := c.uc.StoreUpdateBalances(createOrUpdateBalances)

	if err != nil {
		fmt.Println(err)
	}

	return err
}
