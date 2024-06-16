package entity

import (
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	ID        string     `db:"id,omitempty" validate:"required"`
	AccountID string     `db:"account_id,omitempty" validate:"required"`
	Value     float64    `db:"value,omitempty" validate:"required"`
	CreatedAt *time.Time `db:"created_at,omitempty" validate:"required"`
	UpdatedAt *time.Time `db:"updated_at,omitempty" validate:"required"`
}

func NewBalance(
	accountId string,
	value float64,
) *Balance {
	now := time.Now()

	return &Balance{
		uuid.New().String(),
		accountId,
		value,
		&now,
		&now,
	}
}
