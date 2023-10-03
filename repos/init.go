package repos

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func InitProvinceRepository() ProvinceRepository {
	file, err := os.Open("./data/province.json")
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
	return repos
}

func InitDistrictRepository() DistrictRepos {
	file, err := os.Open("./data/district.json")
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
	return repos
}

func InitWardRepository() WardRepos {
	file, err := os.Open("./data/ward.json")
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

	return repos
}
