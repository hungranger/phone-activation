package cliClient_test

import (
	"phone_activation/pkg/interface/cliClient"
	"phone_activation/pkg/mock"
	"phone_activation/pkg/model"
	"reflect"
	"testing"
	"time"
)

func TestHandler_GetFirstActivationDates(t *testing.T) {
	a1, _ := time.Parse(mock.DATE_FORMAT, "2016-06-01")
	d1, _ := time.Parse(mock.DATE_FORMAT, "2016-09-01")
	a2, _ := time.Parse(mock.DATE_FORMAT, "2016-02-01")
	d2, _ := time.Parse(mock.DATE_FORMAT, "2016-03-01")
	a3, _ := time.Parse(mock.DATE_FORMAT, "2016-01-01")
	d3, _ := time.Parse(mock.DATE_FORMAT, "2016-01-10")
	want := []model.Phone{
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
	}

	tests := []struct {
		name    string
		h       *cliClient.Handler
		want    []model.Phone
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"Get Successfully",
			cliClient.NewHandler(mock.ListActivationDatesUsecaseSuccessStub{}),
			want,
			false,
		},
		{
			"Get Unsuccessfully",
			cliClient.NewHandler(mock.ListActivationDatesUsecaseFailStub{}),
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.GetFirstActivationDates()
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetFirstActivationDates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.GetFirstActivationDates() = %v, want %v", got, tt.want)
			}
		})
	}
}
