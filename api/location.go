package api

import (
	repos2 "delivery-service/domain/repos"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type VietNamProvinceHandle struct {
	provinceRepo *repos2.ProvinceRepository
	districtRepo *repos2.DistrictRepos
	wardRepo     *repos2.WardRepos
}

func NewVietNamProvinceHandle(provinceRepo *repos2.ProvinceRepository,
	districtRepo *repos2.DistrictRepos,
	wardRepo *repos2.WardRepos) *VietNamProvinceHandle {
	return &VietNamProvinceHandle{
		provinceRepo: provinceRepo,
		districtRepo: districtRepo,
		wardRepo:     wardRepo,
	}
}

func (api VietNamProvinceHandle) GetAllProvince(ctx *fiber.Ctx) error {
	dataResp := api.provinceRepo.GetAll()
	return ctx.JSON(dataResp)
}

func (api VietNamProvinceHandle) GetAllDistrictByProvince(ctx *fiber.Ctx) error {
	key := ctx.Params("id")
	if key == "" {
		return ctx.Status(http.StatusBadRequest).SendString("not found province key")
	} else {
		dataResp := api.districtRepo.GetByProvinceKey(key)
		if len(*dataResp) == 0 {
			return ctx.Status(http.StatusNotFound).SendString("not found")
		}
		return ctx.JSON(dataResp)
	}
}

func (api VietNamProvinceHandle) GetAllWardByDistrict(ctx *fiber.Ctx) error {
	key := ctx.Params("id")
	if key == "" {
		return ctx.Status(http.StatusBadRequest).SendString("not found district key")
	} else {
		dataResp := api.wardRepo.GetByDistrictKey(key)
		if len(*dataResp) == 0 {
			return ctx.Status(http.StatusNotFound).SendString("not found")
		}
		return ctx.JSON(dataResp)
	}
}
