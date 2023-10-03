package repos

import "delivery-service/domain"

type DistrictRepos struct {
	Data map[string]domain.DistrictDetail
}

func (repo DistrictRepos) GetByKey(key string) domain.DistrictDetail {
	return repo.Data[key]
}

func (repo DistrictRepos) GetByProvinceKey(provinceKey string) *map[string]domain.DistrictDetail {
	resp := make(map[string]domain.DistrictDetail)
	for key, value := range repo.Data {
		if value.ParentCode == provinceKey {
			resp[key] = value
		}
	}
	return &resp
}

func (repo DistrictRepos) GetAll() map[string]domain.DistrictDetail {
	return repo.Data
}
