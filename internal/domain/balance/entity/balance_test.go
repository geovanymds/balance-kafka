package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewBalance(t *testing.T) {
	balance := NewBalance("any_id", 200, nil, nil)
	assert.Equal(t, "any_id", balance.AccountID)
	assert.Equal(t, 200, balance.Value)
}
