package repos

import (
	"delivery-service/internal/domain/entities"
)

type WardRepos struct {
	Data map[string]entities.WardDetail
}

func (repo WardRepos) GetByKey(key string) entities.WardDetail {
	return repo.Data[key]
}

func (repo WardRepos) GetAll() map[string]entities.WardDetail {
	return repo.Data
}

func (repo WardRepos) GetByDistrictKey(districtKey string) *map[string]entities.WardDetail {
	resp := make(map[string]entities.WardDetail)
	for key, value := range repo.Data {
		if value.ParentCode == districtKey {
			resp[key] = value
		}
	}
	return &resp
}
