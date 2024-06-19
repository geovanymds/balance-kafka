package consumer_proccess_balances

import (
	"encoding/json"
	"fmt"

	"github.com/geovanymds/balance/internal/domain/balance/dto"
)

type MessageContent struct {
	AccountIdFrom      string  `json:"account_id_from" validate:"required"`
	AccountIdTo        string  `json:"account_id_to" validate:"required"`
	BalanceAccountFrom float64 `json:"balance_account_id_from" validate:"required"`
	BalanceAccountTo   float64 `json:"balance_account_id_to" validate:"required"`
}

type Message struct {
	Name    string
	Payload MessageContent
}

func (c *ConsumerBalances) Proccess(msg []byte) error {

	var messageParsed Message

	err := json.Unmarshal(msg, &messageParsed)

	if err != nil {
		fmt.Println(err)
	}

	createOrUpdateBalances := &dto.CreateOrUpdateBalanceInputDto{
		AccountIDFrom:      messageParsed.Payload.AccountIdFrom,
		AccountIDTo:        messageParsed.Payload.AccountIdTo,
		BalanceAccountFrom: messageParsed.Payload.BalanceAccountFrom,
		BalanceAccountTo:   messageParsed.Payload.BalanceAccountTo,
	}

	err = c.uc.StoreUpdateBalances(createOrUpdateBalances)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
