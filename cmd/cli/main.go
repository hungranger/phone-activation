package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"phone_activation/pkg/config"
	"phone_activation/pkg/container"
	"phone_activation/pkg/interface/cliClient"
	"phone_activation/pkg/model"
	"phone_activation/pkg/usecase"
	"runtime"

	"github.com/pkg/errors"
)

const (
	DEV_CONFIG  string = "../../pkg/config/appConfigDev.yaml"
	PROD_CONFIG string = "../../pkg/config/appConfigProd.yaml"
	OUTPUT_PATH string = "output.csv"
)

func main() {
	configPath := DEV_CONFIG
	di, err := loadConfig(configPath)
	if err != nil {
		log.Fatalf("%+v", err)
		return
	}

	inputFile, outputFile := getArgs()
	if inputFile != "" {
		di.(*container.Container).AppConfig.UseCase.ListFirstActivationDates.PhoneRepoConfig.DataStoreConfig.ConnectionString = inputFile
	}

	listActDateUseCase, err := di.BuildUseCase(config.LIST_FIRST_ACTIVATION_DATES)
	if errors.Cause(err) == config.ErrInputFileNotFound {
		log.Printf("Input file '%s' Not Found. 'Err= %v'\n", inputFile, err.Error())
		return
	}
	if err != nil {
		log.Fatalf("%+v", err)
		return
	}

	handler := cliClient.NewHandler(listActDateUseCase.(usecase.IListActivationDatesUsecase))
	firstActivationDates, err := handler.GetFirstActivationDates()
	if errors.Cause(err) == config.ErrDataNotFound {
		log.Printf("There no data in csv file. 'Err= %v'\n", err.Error())
		return
	}
	if errors.Cause(err) == config.ErrDataInvalid {
		log.Printf("Invalid data in csv file. 'Err= %v'\n", err.Error())
		return
	}

	writeToCSV(outputFile, firstActivationDates)

	PrintMemUsage()
}

func loadConfig(filePath string) (container.IContainer, error) {
	factoryMap := make(map[string]interface{})
	appConfig := config.AppConfig{}
	container := container.Container{factoryMap, &appConfig}
	err := container.InitApp(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return &container, nil
}

func getArgs() (inputPath string, outputPath string) {
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	}
	if len(os.Args) > 2 {
		outputPath = os.Args[2]
	}
	return
}

func writeToCSV(filename string, records []model.Phone) {
	if filename == "" {
		filename = OUTPUT_PATH
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	err = w.Write([]string{"PHONE_NUMBER", "REAL_ACTIVATION_DATE"})
	if err != nil {
		log.Fatalf("%+v", err)
	}

	for _, record := range records {
		data := []string{fmt.Sprintf("%s", record.PhoneNumber), record.ActivationDate.Format("2006-01-02")}
		if err := w.Write(data); err != nil {
			log.Fatalf("error writing record to csv: %+v \n", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatalf("%+v", err)
	}
	log.Fatalf("Writing to '%s' successfully\n", filename)
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
