package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	consumer_manager "github.com/geovanymds/balance/events"
	"github.com/geovanymds/balance/internal/domain/balance/usecase"
	web "github.com/geovanymds/balance/internal/infra/api"
	"github.com/geovanymds/balance/internal/infra/api/webserver"
	"github.com/geovanymds/balance/internal/infra/config"
	database "github.com/geovanymds/balance/internal/infra/db"
	"github.com/geovanymds/balance/internal/infra/kafka"
	"github.com/geovanymds/balance/internal/infra/repository"
)

func main() {
	db, err := database.InitDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = database.ExecSeeds(db)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	ctx := context.Background()

	balanceDb := database.NewClientDB(db)
	balanceRepository := repository.NewBalanceRepository(balanceDb)
	balanceUsecase := usecase.NewBalanceUseCase(balanceRepository)

	webserver := webserver.NewWebServer(":3003")

	balanceHandler := web.NewBalanceController(&ctx, balanceUsecase)

	webserver.AddHandler("/balances/{accountId}", balanceHandler.GetBalance)

	fmt.Println("Balance server is running")
	go webserver.Start()

	consumerManager := consumer_manager.InitConsumers(db)

	consumer := kafka.NewConsumer(config.NewKafkaConfig(), []string{"balances"})
	channel := make(chan *ckafka.Message)
	go func() {
		for kafkaMessage := range channel {
			var parsedMessage kafka.Message

			err = json.Unmarshal(kafkaMessage.Value, &parsedMessage)

			if err == nil {
				consumerManager.Consumers[parsedMessage.Name].Proccess(kafkaMessage.Value)
			}
		}
	}()
	consumer.Consume(channel)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan
}
