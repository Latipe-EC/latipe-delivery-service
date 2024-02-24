package api

import (
	_ "delivery-service/docs"
	"delivery-service/internal/domain/repos"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type VietNamProvinceHandle struct {
	provinceRepo *repos.ProvinceRepository
	districtRepo *repos.DistrictRepos
	wardRepo     *repos.WardRepos
}

func NewVietNamProvinceHandle(provinceRepo *repos.ProvinceRepository,
	districtRepo *repos.DistrictRepos,
	wardRepo *repos.WardRepos) *VietNamProvinceHandle {
	return &VietNamProvinceHandle{
		provinceRepo: provinceRepo,
		districtRepo: districtRepo,
		wardRepo:     wardRepo,
	}
}

// @Summary Get all province
// @Tags Location
// @Description Get all province
// @Accept json
// @Produce json
// @Router /vn-location/province [get]
// @Success 200 {array} Province
// @Failure 400 {object} DefaultResponse
func (api VietNamProvinceHandle) GetAllProvince(ctx *fiber.Ctx) error {
	dataResp := api.provinceRepo.GetAll()
	return ctx.JSON(dataResp)
}

// @Summary Get all district by province
// @Tags Location
// @Description Get all district by province
// @Accept json
// @Produce json
// @Param id path string true "province id"
// @Router /vn-location/district/{id} [get]
// @Success 200 {array} District
// @Failure 400 {object} DefaultResponse
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

// @Summary Get all ward by district
// @Tags Location
// @Description Get all ward by district
// @Accept json
// @Produce json
// @Param id path string true "district id"
// @Router /vn-location/ward/{id} [get]
// @Success 200 {array} Ward
// @Failure 400 {object} DefaultResponse
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
