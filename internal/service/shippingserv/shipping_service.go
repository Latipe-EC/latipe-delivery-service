package shippingserv

import (
	"context"
	"delivery-service/internal/adapter/userserv"
	"delivery-service/internal/domain/dto"
	"delivery-service/internal/domain/entities"
	"delivery-service/internal/domain/repos"
	"errors"
	"time"
)

type ShippingCostService struct {
	provinceRepo *repos.ProvinceRepository
	userService  *userserv.UserService
	deliRepo     *repos.DeliveryRepos
}

func NewShippingCostService(provinceRepo *repos.ProvinceRepository, userService *userserv.UserService,
	deliveryRepos *repos.DeliveryRepos) *ShippingCostService {
	return &ShippingCostService{
		provinceRepo: provinceRepo,
		userService:  userService,
		deliRepo:     deliveryRepos,
	}
}

func (sh ShippingCostService) CalculateByProvinceCode(ctx context.Context,
	req *dto.CalculateShippingCostRequest) ([]*dto.CalculateShippingCostShipping, error) {

	var storeLocation []entities.ProvinceDetail
	for _, i := range req.SrcCode {
		src := sh.provinceRepo.GetByKey(i)
		storeLocation = append(storeLocation, src)
	}

	dest := sh.provinceRepo.GetByKey(req.DestCode)

	deliveries, err := sh.deliRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(req.SrcCode) < 1 || dest.Code == "" || len(deliveries) == 0 {
		return nil, errors.New("not found")
	}

	var resp []*dto.CalculateShippingCostShipping
	for _, deli := range deliveries {
		if deli.IsActive != false {
			cost := 0
			var receive time.Time
			for _, i := range storeLocation {
				c, r := CalculateShippingCodes(i.Code, dest.Code, deli.BaseCost)
				cost += c
				receive = r
			}

			layout := "2006-01-02"
			formattedTime := receive.Format(layout)

			data := dto.CalculateShippingCostShipping{
				SrcCode:      req.SrcCode,
				DestCode:     dest.Code,
				ReceiveDate:  formattedTime,
				DeliveryId:   deli.ID.Hex(),
				DeliveryName: deli.DeliveryName,
				Cost:         cost,
			}
			resp = append(resp, &data)
		}
	}

	return resp, err
}

func (sh ShippingCostService) CalculateOrderShippingCost(ctx context.Context,
	req *dto.OrderShippingCostRequest) (*dto.CalculateShippingCostShipping, error) {

	src := sh.provinceRepo.GetByKey(req.SrcCode)
	dest := sh.provinceRepo.GetByKey(req.DestCode)

	delivery, err := sh.deliRepo.GetById(ctx, req.DeliveryId)
	if err != nil {
		return nil, err
	}

	if delivery == nil || delivery.IsActive == false {
		return nil, errors.New("not found")
	}

	data := dto.CalculateShippingCostShipping{
		DeliveryId:   delivery.ID.Hex(),
		DeliveryName: delivery.DeliveryName,
	}

	c, receive := CalculateShippingCodes(dest.Code, src.Code, delivery.BaseCost)
	formattedTime := receive.Format("2006-01-02")
	data.ReceiveDate = formattedTime

	data.Cost = c

	return &data, err
}
