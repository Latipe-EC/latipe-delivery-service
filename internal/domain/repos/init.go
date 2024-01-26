package repos

import (
	"delivery-service/config"
	"encoding/json"
	"io"
	"log"
	"os"
)

func InitProvinceRepository(config *config.Config) *ProvinceRepository {
	file, err := os.Open(config.VietNamLocationData.ProvincePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	byteVal, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var repos ProvinceRepository

	err = json.Unmarshal(byteVal, &repos.Data)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	log.Printf("[%s] Init province repos was successful", "Info")
	return &repos
}

func InitDistrictRepository(config *config.Config) *DistrictRepos {
	file, err := os.Open(config.VietNamLocationData.DistrictPath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	byteVal, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var repos DistrictRepos

	err = json.Unmarshal(byteVal, &repos.Data)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	log.Printf("[%s] Init district repos was successful", "Info")
	return &repos
}

func InitWardRepository(config *config.Config) *WardRepos {
	file, err := os.Open(config.VietNamLocationData.WardPath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	byteVal, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var repos WardRepos

	err = json.Unmarshal(byteVal, &repos.Data)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	log.Printf("[%s] Init ward repos was successful", "Info")

	return &repos
}
