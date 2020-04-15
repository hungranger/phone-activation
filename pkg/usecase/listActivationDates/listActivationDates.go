package listActivationDates

import (
	"phone_activation/pkg/model"
	"phone_activation/pkg/repository"
)

type ListActivationDatesUsecase struct {
	phoneRepo repository.IPhoneRepository
}

func NewListActivationDatesUsecase(phoneRepo repository.IPhoneRepository) *ListActivationDatesUsecase {
	return &ListActivationDatesUsecase{phoneRepo}
}

func (uc *ListActivationDatesUsecase) FindAll() ([]model.Phone, error) {
	return uc.phoneRepo.FindAll()
}
