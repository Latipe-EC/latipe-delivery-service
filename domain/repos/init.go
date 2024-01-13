package repos

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func InitProvinceRepository(opt bool, optPath ...string) ProvinceRepository {

	path := "./data/vn_data/province.json" //default
	if len(optPath) > 0 && opt == true {
		path = optPath[0]
	}

	file, err := os.Open(path)
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

func InitDistrictRepository(opt bool, optPath ...string) DistrictRepos {
	path := "/data/vn_data/district.json" //default
	if len(optPath) > 0 && opt == true {
		path = optPath[0]
	}

	file, err := os.Open(path)
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

func InitWardRepository(opt bool, optPath ...string) WardRepos {
	path := "./data/vn_data/ward.json" //default
	if len(optPath) > 0 && opt == true {
		path = optPath[0]
	}

	file, err := os.Open(path)
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
