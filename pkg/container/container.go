package container

import (
	"phone_activation/pkg/config"
	"phone_activation/pkg/repository"
	"phone_activation/pkg/repository/PhoneRepository/csvdb"
	"phone_activation/pkg/usecase/listActivationDates"

	"github.com/pkg/errors"
)

type IContainer interface {
	// InitApp loads the application configurations from a file and saved it in appConfig and initialize the logger
	// The appConfig is cached in container, so it only loads the configuration file once.
	// InitApp only needs to be called once. If the configuration changes, you can call it again to reinitialize the app.
	InitApp(filename string) error

	// BuildUseCase creates concrete types for use case and it is included types.
	// For each call, it will create a new instance, which means it is not a singleton
	// Only exceptions are data store handlers, which are singletons. They are cached in container.
	BuildUseCase(code string) (interface{}, error)
}

type Container struct {
	FactoryMap map[string]interface{}
	AppConfig  *config.AppConfig
}

func (sc *Container) InitApp(filename string) error {
	var err error
	config, err := loadConfig(filename)
	if err != nil {
		return errors.Wrap(err, "loadConfig")
	}
	sc.AppConfig = config
	return nil
}

func loadConfig(filename string) (*config.AppConfig, error) {
	ac, err := config.ReadConfig(filename)
	if err != nil {
		return nil, errors.Wrap(err, "readConfigFile")
	}
	return ac, nil
}

func (sc *Container) BuildUseCase(code string) (interface{}, error) {
	switch code {
	case config.LIST_FIRST_ACTIVATION_DATES:
		listActDatesCfg := sc.AppConfig.UseCase.ListFirstActivationDates
		phoneRepo, err := sc.buildRepo(&listActDatesCfg.PhoneRepoConfig)
		if err != nil {
			return nil, err
		}
		return listActivationDates.NewListActivationDatesUsecase(phoneRepo), nil
	}
	return nil, nil
}

func (sc *Container) buildRepo(repoCfg *config.RepoConfig) (repository.IPhoneRepository, error) {
	switch repoCfg.Code {
	case config.PHONE_REPO:
		var repo repository.IPhoneRepository
		var err error

		switch repoCfg.DataStoreConfig.Code {
		case config.CSVDB:
			repo, err = csvdb.NewCSVRepository(repoCfg.DataStoreConfig.ConnectionString)
		}

		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		return repo, nil
	}
	return nil, nil
}
