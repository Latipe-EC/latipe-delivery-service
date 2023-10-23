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
	req *dto.CalculateShippingCostRequest) ([]*dto.CalculateShippingCostShipping, error) {

	src := sh.provinceRepo.GetByKey(req.SrcCode)
	dest := sh.provinceRepo.GetByKey(req.DestCode)

	deliveries, err := sh.deliRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if src.Code == "" || dest.Code == "" || len(deliveries) == 0 {
		return nil, errors.New("not found")
	}

	var resp []*dto.CalculateShippingCostShipping
	for _, deli := range deliveries {
		cost, receive := CalculateShippingCodes(src.Code, dest.Code, deli.BaseCost)
		layout := "2006-01-02"
		formattedTime := receive.Format(layout)

		data := dto.CalculateShippingCostShipping{
			SrcCode:      src.Code,
			DestCode:     dest.Code,
			ReceiveDate:  formattedTime,
			DeliveryId:   deli.ID.Hex(),
			DeliveryName: deli.DeliveryName,
			Cost:         cost,
		}

		resp = append(resp, &data)
	}

	return resp, err
}
