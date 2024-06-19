package consumer

import (
	"database/sql"

	"github.com/geovanymds/balance/events/consumer_proccess_balances"
	"github.com/geovanymds/balance/internal/domain/balance/usecase"
	database "github.com/geovanymds/balance/internal/infra/db"
	"github.com/geovanymds/balance/internal/infra/repository"
)

func InitConsumers(db *sql.DB) *ConsumerManager {
	balanceDb := database.NewClientDB(db)
	balanceRepository := repository.NewBalanceRepository(balanceDb)
	balanceUsecase := usecase.NewBalanceUseCase(balanceRepository)
	cm := NewConsumerManager()

	cm.AddConsumer("BalanceUpdated", consumer_proccess_balances.NewConsumerBalances(balanceUsecase))

	return cm
}
