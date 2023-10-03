package repos

import "delivery-service/domain"

type ProvinceRepository struct {
	Data map[string]domain.ProvinceDetail
}

func (repo ProvinceRepository) GetByKey(key string) domain.ProvinceDetail {
	return repo.Data[key]
}

func (repo ProvinceRepository) GetAll() map[string]domain.ProvinceDetail {
	return repo.Data
}
