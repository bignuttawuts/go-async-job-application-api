package usecases

import "fmt"

type CreateApplicationUsecase struct {
}

func NewCreateApplicationUsecase() *CreateApplicationUsecase {
	return &CreateApplicationUsecase{}
}

func (u *CreateApplicationUsecase) Execute() error {
	fmt.Println("Executing CreateApplicationUsecase...")
	return nil
}
