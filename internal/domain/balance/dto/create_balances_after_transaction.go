package dto

type CreateOrUpdateBalanceInputDto struct {
	AccountIDFrom      string  `db:"account_id_from,omitempty" validate:"required"`
	AccountIDTo        string  `db:"account_id_to,omitempty" validate:"required"`
	BalanceAccountFrom float64 `db:"account_from,omitempty" validate:"required"`
	BalanceAccountTo   float64 `db:"account_to,omitempty" validate:"required"`
}

type CreateBalanceOutputDto struct {
	ID string `db:"id,omitempty" validate:"required"`
}
