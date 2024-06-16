package dto

type GetAccountBalanceDto struct {
	AccountId string `json:"account_id" validade:"required"`
}
