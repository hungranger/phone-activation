package csvdb_test

import (
	"phone_activation/pkg/mock"
	"phone_activation/pkg/model"
	"phone_activation/pkg/repository/PhoneRepository/csvdb"
	"reflect"
	"testing"
	"time"
)

func Test_csvRepository_FindAll(t *testing.T) {
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
		name     string
		filePath string
		want     []model.Phone
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			"Get all first activation dates success",
			"valid_phone_test.csv",
			want,
			false,
		},
		{
			"Fail: invalid content format",
			"invalid_phone_test.csv",
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := csvdb.NewCSVRepository(tt.filePath)
			if err != nil {
				t.Errorf("csvRepository.NewCSVRepository() error = %v", err)
				return
			}

			got, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("csvRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("csvRepository.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
