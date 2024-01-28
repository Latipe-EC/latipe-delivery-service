package deliveryserv

import (
	"context"
	"delivery-service/internal/adapter/userserv"
	"delivery-service/internal/adapter/userserv/dto"
	dto2 "delivery-service/internal/domain/dto"
	"delivery-service/internal/domain/entities"
	"delivery-service/internal/domain/repos"
	"delivery-service/pkgs/mapper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeliveryService struct {
	userService *userserv.UserService
	deliRepo    *repos.DeliveryRepos
}

func NewDeliveryService(userService *userserv.UserService,
	deliveryRepos *repos.DeliveryRepos) *DeliveryService {
	return &DeliveryService{
		userService: userService,
		deliRepo:    deliveryRepos,
	}
}

func (dl DeliveryService) GetAllDeliveries(ctx context.Context) (*dto2.GetAllDeliveries, error) {

	deliveries, err := dl.deliRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var resp dto2.GetAllDeliveries

	if err := mapper.BindingStruct(deliveries, &resp.Items); err != nil {
		return nil, err
	}

	return &resp, err
}

func (dl DeliveryService) GetById(ctx context.Context, id string) (*dto2.DeliveryDetail, error) {

	deliveries, err := dl.deliRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	var resp dto2.DeliveryDetail

	if err := mapper.BindingStruct(deliveries, &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

func (dl DeliveryService) GetByUserId(ctx context.Context, userId string) (*dto2.DeliveryDetail, error) {

	deliveries, err := dl.deliRepo.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	var resp dto2.DeliveryDetail

	if err := mapper.BindingStruct(deliveries, &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

func (dl DeliveryService) CreateDelivery(ctx context.Context, deli *dto2.CreateDeliveryRequest) (string, error) {

	req := dto.CreateAccountRequest{
		AuthorizationHeader: dto.AuthorizationHeader{BearerToken: deli.BearerToken},
	}
	req.Body.FirstName = deli.DeliveryName
	req.Body.LastName = "Đơn vị vận chuyển"
	req.Body.PhoneNumber = deli.PhoneNumber
	req.Body.Email = deli.Email
	req.Body.Role = dto.DELIVERY_ROLE

	resp, err := dl.userService.CreateNewAccount(ctx, &req)
	if err != nil {
		return "", err
	}

	entity := entities.Delivery{
		DeliveryName: deli.DeliveryName,
		DeliveryCode: deli.DeliveryCode,
		BaseCost:     deli.BaseCost,
		Description:  deli.Description,
		OwnerAccount: entities.OwnerAccount{
			UserID:      resp.Id,
			PhoneNumber: resp.PhoneNumber,
			Email:       resp.Email,
		},
		IsActive: true,
	}

	deliId, err := dl.deliRepo.CreateDelivery(ctx, &entity)

	if err != nil {
		return "", err
	}

	return deliId, err
}

func (dl DeliveryService) UpdateDelivery(ctx context.Context, deli *dto2.UpdateDeliveryRequest) error {
	id, err := primitive.ObjectIDFromHex(deli.DeliId)
	if err != nil {
		return err
	}

	entity := entities.Delivery{
		ID:           id,
		DeliveryName: deli.DeliveryName,
		DeliveryCode: deli.DeliveryCode,
		BaseCost:     deli.BaseCost,
		Description:  deli.Description,
	}
	err = dl.deliRepo.Update(ctx, &entity)
	if err != nil {
		return err
	}

	return nil
}

func (dl DeliveryService) UpdateStatusDelivery(ctx context.Context, deli *dto2.UpdateStatusDeliveryRequest) error {
	id, err := primitive.ObjectIDFromHex(deli.DeliId)
	if err != nil {
		return err
	}

	entity := entities.Delivery{
		ID:       id,
		IsActive: deli.Status,
	}

	err = dl.deliRepo.UpdateStatus(ctx, &entity)
	if err != nil {
		return err
	}

	return nil
}
