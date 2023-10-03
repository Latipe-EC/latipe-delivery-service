package repos

import "delivery-service/domain"

type WardRepos struct {
	Data map[string]domain.WardDetail
}

func (repo WardRepos) GetByKey(key string) domain.WardDetail {
	return repo.Data[key]
}

func (repo WardRepos) GetAll() map[string]domain.WardDetail {
	return repo.Data
}

func (repo WardRepos) GetByDistrictKey(districtKey string) *map[string]domain.WardDetail {
	resp := make(map[string]domain.WardDetail)
	for key, value := range repo.Data {
		if value.ParentCode == districtKey {
			resp[key] = value
		}
	}
	return &resp
}
