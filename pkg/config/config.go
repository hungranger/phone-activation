package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var (
	ErrDataNotFound      = errors.New("Data Not Found")
	ErrDataInvalid       = errors.New("Data Invalid")
	ErrInputFileNotFound = errors.New("Input File Not Found")
)

type AppConfig struct {
	CSVDBConfig DataStoreConfig `yaml:"csvDBConfig"`
	UseCase     UseCaseConfig   `yaml:"useCaseConfig"`
}

type DataStoreConfig struct {
	Code string `yaml:"code"`
	// Only database has a driver name, for grpc it is "tcp" ( network) for server
	DriverName string `yaml:"driverName"`
	// For csv, this is filepath .For database, this is datasource name; for grpc, it is target url;
	ConnectionString string `yaml:"connectionString"`
	// Only some databases need this password
	Password string `yaml:"password"`
	// Only some databases need this database name
	DbName string `yaml:"dbName"`
	// Only some databases need this timeout
	Timeout int `yaml:"timeout"`
}

type UseCaseConfig struct {
	ListFirstActivationDates ListFirstActivationDatesConfig `yaml:"listFirstActivationDates"`
}

type ListFirstActivationDatesConfig struct {
	Code            string     `yaml:"code"`
	PhoneRepoConfig RepoConfig `yaml:"phoneRepoConfig"`
}

type RepoConfig struct {
	Code            string          `yaml:"code"`
	DataStoreConfig DataStoreConfig `yaml:"dataStoreConfig"`
}

func ReadConfig(filename string) (*AppConfig, error) {
	var ac AppConfig
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read config file error")
	}

	err = yaml.Unmarshal(file, &ac)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}

	err = validateConfig(ac)
	if err != nil {
		return nil, errors.Wrap(err, "validate config")
	}

	return &ac, nil
}
