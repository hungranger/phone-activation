package mock

import (
	"errors"
	"phone_activation/pkg/model"
)

type ListActivationDatesUsecaseFailStub struct {
}

func (r ListActivationDatesUsecaseFailStub) FindAll() ([]model.Phone, error) {
	return nil, errors.New("Can not read file csv")
}
