package usecase

import (
	"github.com/geovanymds/balance/internal/infra/repository"
)

type BalanceUseCase struct {
	Repository *repository.BalanceRepository
}

func NewBalanceUseCase(r *repository.BalanceRepository) *BalanceUseCase {
	return &BalanceUseCase{r}
}
