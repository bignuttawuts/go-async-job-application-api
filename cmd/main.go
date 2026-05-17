package main

import (
	"github.com/bignuttawuts/go-async-job-application-api/config"
	"github.com/bignuttawuts/go-async-job-application-api/internal/adapters/in/rest"
	"github.com/bignuttawuts/go-async-job-application-api/internal/adapters/out/publisher"
	"github.com/bignuttawuts/go-async-job-application-api/internal/usecases"
)

func main() {
	cfg := config.NewConfig()
	pub := publisher.NewPublisher(cfg.Kafka.Brokers, cfg.Producer.Topic)

	ucs := &rest.UseCases{
		CreateApplicationUsecase: usecases.NewCreateApplicationUsecase(pub),
	}

	srv := rest.NewServer(ucs, cfg)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
