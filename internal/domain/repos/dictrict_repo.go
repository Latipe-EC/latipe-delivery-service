package repos

import (
	"delivery-service/internal/domain/entities"
)

type DistrictRepos struct {
	Data map[string]entities.DistrictDetail
}

func (repo DistrictRepos) GetByKey(key string) entities.DistrictDetail {
	return repo.Data[key]
}

func (repo DistrictRepos) GetByProvinceKey(provinceKey string) *map[string]entities.DistrictDetail {
	resp := make(map[string]entities.DistrictDetail)
	for key, value := range repo.Data {
		if value.ParentCode == provinceKey {
			resp[key] = value
		}
	}
	return &resp
}

func (repo DistrictRepos) GetAll() map[string]entities.DistrictDetail {
	return repo.Data
}
