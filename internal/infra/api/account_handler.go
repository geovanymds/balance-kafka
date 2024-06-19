package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/geovanymds/balance/internal/domain/balance/dto"
	"github.com/geovanymds/balance/internal/domain/balance/usecase"
)

type BalanceController struct {
	BalanceUsecase *usecase.BalanceUseCase
}

func NewBalanceController(ctx *context.Context, uc *usecase.BalanceUseCase) *BalanceController {
	return &BalanceController{
		BalanceUsecase: uc,
	}
}

func (h *BalanceController) GetBalance(w http.ResponseWriter, r *http.Request) {
	accountID := strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1]

	if accountID == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Println("Account ID not provided")
		return
	}

	output, err := h.BalanceUsecase.GetBalanceByAccountId(&dto.GetAccountBalanceDto{AccountId: accountID})

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Println("Account ID not provided")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
