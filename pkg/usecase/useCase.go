package usecase

import "phone_activation/pkg/model"

type IListActivationDatesUsecase interface {
	FindAll() ([]model.Phone, error)
}
