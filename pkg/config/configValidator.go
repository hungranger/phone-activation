package config

import (
	"github.com/pkg/errors"
)

// database code. Need to map to the database code (DataStoreConfig) in the configuration yaml file.
const (
	MONGODB string = "mongodb"
	CSVDB   string = "csvdb"
)

// use case code. Need to map to the use case code (UseCaseConfig) in the configuration yaml file.
// Client app use those to retrieve use case from the container
const (
	LIST_FIRST_ACTIVATION_DATES string = "listFirstActivationDates"
)

const (
	PHONE_REPO string = "phoneRepo"
)

func validateConfig(appConfig AppConfig) error {
	err := validateDataStore(appConfig)
	if err != nil {
		return err
	}

	useCase := appConfig.UseCase
	err = validateUseCase(useCase)
	if err != nil {
		return err
	}
	return nil
}

func validateDataStore(appConfig AppConfig) error {
	mgc := appConfig.CSVDBConfig
	key := mgc.Code
	mgcMsg := " in validateDataStore doesn't match key = "
	if CSVDB != key {
		errMsg := CSVDB + mgcMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateUseCase(useCase UseCaseConfig) error {
	err := validateListFirstActivationDates(useCase)
	if err != nil {
		return err
	}
	return nil
}

func validateListFirstActivationDates(usecase UseCaseConfig) error {
	sc := usecase.ListFirstActivationDates
	key := sc.Code
	scMsg := " in validateShortenURL doesn't match key = "
	if LIST_FIRST_ACTIVATION_DATES != key {
		errMsg := LIST_FIRST_ACTIVATION_DATES + scMsg + key
		return errors.New(errMsg)
	}

	key = sc.PhoneRepoConfig.Code
	if PHONE_REPO != key {
		errMsg := PHONE_REPO + scMsg + key
		return errors.New(errMsg)
	}
	return nil
}
