package entity

import (
	"time"
)

type Balance struct {
	AccountID string     `db:"account_id,omitempty" validate:"required"`
	Value     float64    `db:"value,omitempty" validate:"required"`
	CreatedAt *time.Time `db:"created_at,omitempty" validate:"required"`
	UpdatedAt *time.Time `db:"updated_at,omitempty" validate:"required"`
}

func NewBalance(
	accountId string,
	value float64,
	createdAt *time.Time,
	updatedAt *time.Time,

) *Balance {
	if createdAt == nil {
		now := time.Now()
		return &Balance{
			accountId,
			value,
			&now,
			&now,
		}
	}
	return &Balance{
		accountId,
		value,
		createdAt,
		updatedAt,
	}
}
