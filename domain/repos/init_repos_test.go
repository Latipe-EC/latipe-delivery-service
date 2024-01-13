package repos

import (
	"delivery-service/domain/entities"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestInitDistrictRepository(t *testing.T) {
	repos := InitDistrictRepository(true, "../../data/vn_data/district.json")

	data := repos.GetByKey("001")
	expectValue := entities.DistrictDetail{
		Name:         "Ba Đình",
		Type:         "quan",
		Slug:         "ba-dinh",
		NameWithType: "Quận Ba Đình",
		Path:         "Ba Đình, Hà Nội",
		PathWithType: "Quận Ba Đình, Thành phố Hà Nội",
		Code:         "001",
		ParentCode:   "01",
	}

	t.Run("init district repository and get key", func(t *testing.T) {
		assert.Equal(t, data, expectValue)
	})
}

func TestInitProvinceRepository(t *testing.T) {
	repos := InitProvinceRepository(true, "../../data/vn_data/province.json")

	data := repos.GetByKey("01")
	expectValue := entities.ProvinceDetail{
		Name:         "Hà Nội",
		Slug:         "ha-noi",
		Type:         "thanh-pho",
		NameWithType: "Thành phố Hà Nội",
		Code:         "01",
	}

	t.Run("init province repository and get key", func(t *testing.T) {
		assert.Equal(t, data, expectValue)
	})
}

func TestInitWardRepository(t *testing.T) {
	repos := InitWardRepository(true, "../../data/vn_data/ward.json")

	data := repos.GetByKey("00001")
	expectValue := entities.WardDetail{
		Name:         "Phúc Xá",
		Type:         "phuong",
		Slug:         "phuc-xa",
		NameWithType: "Phường Phúc Xá",
		Path:         "Phúc Xá, Ba Đình, Hà Nội",
		PathWithType: "Phường Phúc Xá, Quận Ba Đình, Thành phố Hà Nội",
		Code:         "00001",
		ParentCode:   "001",
	}

	t.Run("init ward repository and get key", func(t *testing.T) {
		assert.Equal(t, data, expectValue)
	})
}
