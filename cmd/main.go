package main

import (
	"time"

	"github.com/bignuttawuts/go-async-job-application-api/config"
	"github.com/bignuttawuts/go-async-job-application-api/internal/adapters/in/rest"
)

func main() {
	srv := rest.NewServer(config.Config{
		Port:         ":3000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	})
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
