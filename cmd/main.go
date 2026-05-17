package main

import (
	"time"

	"github.com/bignuttawuts/go-async-job-application-api/config"
	"github.com/bignuttawuts/go-async-job-application-api/internal/adapters/in/rest"
	"github.com/bignuttawuts/go-async-job-application-api/internal/adapters/out/publisher"
	"github.com/bignuttawuts/go-async-job-application-api/internal/usecases"
)

func main() {

	pub := publisher.NewPublisher([]string{"localhost:9092"}, "fct.recruitment.job-applications.v1")

	ucs := &rest.UseCases{
		CreateApplicationUsecase: usecases.NewCreateApplicationUsecase(pub),
	}

	srv := rest.NewServer(ucs, config.Config{
		Port:         ":3000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	})
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
