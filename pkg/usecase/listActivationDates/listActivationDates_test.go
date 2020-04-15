package listActivationDates_test

import (
	"phone_activation/pkg/mock"
	"phone_activation/pkg/model"
	"phone_activation/pkg/usecase/listActivationDates"
	"reflect"
	"testing"
	"time"
)

func TestListActivationDatesUsecase_FindAll(t *testing.T) {
	uc := listActivationDates.NewListActivationDatesUsecase(mock.PhoneReposiotySuccessStub{})
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
		uc      *listActivationDates.ListActivationDatesUsecase
		want    []model.Phone
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"Get Successfully",
			uc,
			want,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListActivationDatesUsecase.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListActivationDatesUsecase.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
