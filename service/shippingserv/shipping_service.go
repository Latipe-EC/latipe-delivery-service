package shippingserv

import (
	"context"
	"delivery-service/adapter/userserv"
	"delivery-service/domain/dto"
	"delivery-service/domain/repos"
	"errors"
)

type ShippingCostService struct {
	provinceRepo *repos.ProvinceRepository
	userService  *userserv.UserService
	deliRepo     *repos.DeliveryRepos
}

func NewShippingCostService(provinceRepo *repos.ProvinceRepository, userService *userserv.UserService,
	deliveryRepos *repos.DeliveryRepos) ShippingCostService {
	return ShippingCostService{
		provinceRepo: provinceRepo,
		userService:  userService,
		deliRepo:     deliveryRepos,
	}
}

func (sh ShippingCostService) CalculateByProvinceCode(ctx context.Context,
	req *dto.CalculateShippingCostRequest) (*dto.CalculateShippingCostShipping, error) {

	src := sh.provinceRepo.GetByKey(req.SrcCode)
	dest := sh.provinceRepo.GetByKey(req.DestCode)

	delivery, err := sh.deliRepo.GetById(ctx, req.DeliveryId)
	if err != nil {
		return nil, err
	}

	if src.Code == "" || dest.Code == "" || delivery == nil {
		return nil, errors.New("not found")
	}

	cost, receive := CalculateShippingCodes(src.Code, dest.Code, delivery.BaseCost)
	layout := "2006-01-02"
	formattedTime := receive.Format(layout)

	resp := dto.CalculateShippingCostShipping{
		SrcCode:      src.Code,
		DestCode:     dest.Code,
		ReceiveDate:  formattedTime,
		DeliveryId:   delivery.ID.Hex(),
		DeliveryName: delivery.DeliveryName,
		Cost:         cost,
	}

	return &resp, err
}
