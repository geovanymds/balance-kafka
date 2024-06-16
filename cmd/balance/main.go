package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
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

	webserver := webserver.NewWebServer(":8000")

	balanceHandler := web.NewBalanceController(&ctx, balanceUsecase)

	webserver.AddHandler("/balances/{accountId}", balanceHandler.GetBalance)

	consumer := kafka.NewConsumer(config.NewKafkaConfig(), []string{"balances"})
	channel := make(chan *ckafka.Message)

	// balanceUpdatedConsumer := consumer_proccess_balances.NewConsumerBalances(balanceUsecase)

	go func() {
		myMessage := <-channel

		fmt.Printf("Message on %s: %s\n", *myMessage.TopicPartition.Topic, string(myMessage.Value))
		// if myMessage.Value == "BalanceUpdated" {
		// 	balanceUpdatedConsumer.Proccess()
		// }
		consumer.Consume(channel)
		fmt.Println("Proccessing queue")
	}()

	go func() {
		webserver.Start()
		fmt.Println("Balance server is running")
	}()

	fmt.Println("Server and consumer started. To exit press CTRL+C")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Println("Shutting down...")

}
