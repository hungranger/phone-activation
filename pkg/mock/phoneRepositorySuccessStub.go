package mock

import (
	"phone_activation/pkg/model"
	"time"
)

const (
	DATE_FORMAT string = "2006-01-02"
)

type PhoneReposiotySuccessStub struct {
}

func (r PhoneReposiotySuccessStub) FindAll() ([]model.Phone, error) {
	a1, _ := time.Parse(DATE_FORMAT, "2016-06-01")
	d1, _ := time.Parse(DATE_FORMAT, "2016-09-01")
	a2, _ := time.Parse(DATE_FORMAT, "2016-02-01")
	d2, _ := time.Parse(DATE_FORMAT, "2016-03-01")
	a3, _ := time.Parse(DATE_FORMAT, "2016-01-01")
	d3, _ := time.Parse(DATE_FORMAT, "2016-01-10")
	return []model.Phone{
		model.Phone{
			PhoneNumber:      "0987000001",
			ActivationDate:   a1,
			DeactivationDate: d1,
		},
		model.Phone{
			PhoneNumber:      "0987000002",
			ActivationDate:   a2,
			DeactivationDate: d2,
		},
		model.Phone{
			PhoneNumber:      "0987000003",
			ActivationDate:   a3,
			DeactivationDate: d3,
		},
	}, nil
}
