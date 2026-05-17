package usecases

import (
	"fmt"

	"github.com/bignuttawuts/go-async-job-application-api/internal/adapters/out/publisher"
)

type CreateApplicationUsecase struct {
	pub *publisher.Publisher
}

func NewCreateApplicationUsecase(pub *publisher.Publisher) *CreateApplicationUsecase {
	return &CreateApplicationUsecase{
		pub: pub,
	}
}

func (u *CreateApplicationUsecase) Execute() error {
	fmt.Println("Executing CreateApplicationUsecase...")
	err := u.pub.Publish([]byte("New application received"))
	if err != nil {
		fmt.Printf("Failed to publish message: %v\n", err)
		return err
	}
	return nil
}
