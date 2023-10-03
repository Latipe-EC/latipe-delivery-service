package repos

import (
	"delivery-service/domain/entities"
)

type ProvinceRepository struct {
	Data map[string]entities.ProvinceDetail
}

func (repo ProvinceRepository) GetByKey(key string) entities.ProvinceDetail {
	return repo.Data[key]
}

func (repo ProvinceRepository) GetAll() map[string]entities.ProvinceDetail {
	return repo.Data
}
