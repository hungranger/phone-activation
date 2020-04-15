package mock

import (
	"errors"
	"phone_activation/pkg/model"
)

type PhoneReposiotyFailStub struct {
}

func (r PhoneReposiotyFailStub) FindAll() ([]model.Phone, error) {
	return nil, errors.New("Can not read file csv")
}
