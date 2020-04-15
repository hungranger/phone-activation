package csvdb

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"phone_activation/pkg/config"
	"phone_activation/pkg/model"
	"phone_activation/pkg/repository"
	"sort"
	"time"

	"github.com/pkg/errors"
)

const (
	DATE_FORMAT string = "2006-01-02"
)

type csvRepository struct {
	reader *csv.Reader
	closer io.Closer
}

func newCSVReader(filepath string) (*csv.Reader, io.Closer, error) {
	// Open CSV file
	f, err := os.Open(filepath)
	if err != nil {
		return nil, nil, errors.Wrap(err, "repo.newCSVReader")
	}

	// Read File into a Variable
	csvReader := csv.NewReader(f)
	return csvReader, f, nil
}

func NewCSVRepository(filepath string) (repository.IPhoneRepository, error) {
	// if len(os.Args) > 1 {
	// 	filepath = os.Args[1]
	// }
	// log.Println(filepath)
	csvReader, closer, err := newCSVReader(filepath)
	if err != nil {
		return nil, errors.Wrap(err, "repo.NewCSVRepository")
	}
	log.Printf("Open file '%s' Successfully, ready to read.\n", filepath)
	return &csvRepository{csvReader, closer}, nil
}

func (repo csvRepository) FindAll() ([]model.Phone, error) {
	defer repo.closer.Close()

	// group all records by phone number
	m, err := groupPhoneNumber(repo.reader)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	outputs := getFirstActivationDates(m)
	return outputs, nil
}

func groupPhoneNumber(reader *csv.Reader) (map[string][]model.Phone, error) {
	// remove header
	_, err := reader.Read()
	if err != nil {
		return nil, config.ErrDataNotFound
	}

	m := make(map[string][]model.Phone)
	// stream line by line & turn into object
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, config.ErrDataInvalid
		}

		phoneNumber := line[0]
		activationDate, _ := time.Parse(DATE_FORMAT, line[1])
		deactivationDate, _ := time.Parse(DATE_FORMAT, line[2])

		data := model.Phone{
			PhoneNumber:      phoneNumber,
			ActivationDate:   activationDate,
			DeactivationDate: deactivationDate,
		}

		m[phoneNumber] = append(m[phoneNumber], data)
	}
	return m, nil
}

func getFirstActivationDates(m map[string][]model.Phone) (outputs []model.Phone) {
	for _, records := range m {
		sort.SliceStable(records, func(a, b int) bool {
			return records[a].ActivationDate.Before(records[b].ActivationDate)
		})

		firstActivation := records[0]
		for i := 1; i < len(records); i++ {
			if !records[i].ActivationDate.Equal(records[i-1].DeactivationDate) {
				firstActivation = records[i]
			}
		}

		outputs = append(outputs, firstActivation)
		// delete(m, key)
		log.Printf("%s %s\n", firstActivation.PhoneNumber, firstActivation.ActivationDate.Format(DATE_FORMAT))
	}
	// log.Printf("%v\n", outputs)
	return
}
