package repository

import "phone_activation/pkg/model"

type IPhoneRepository interface {
	FindAll() ([]model.Phone, error)
}
