package deliveryserv

import (
	"context"
	"delivery-service/adapter/userserv"
	"delivery-service/domain/dto"
	"delivery-service/domain/entities"
	"delivery-service/domain/repos"
	"delivery-service/mapper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeliveryService struct {
	userService *userserv.UserService
	deliRepo    *repos.DeliveryRepos
}

func NewDeliveryService(userService *userserv.UserService,
	deliveryRepos *repos.DeliveryRepos) DeliveryService {
	return DeliveryService{
		userService: userService,
		deliRepo:    deliveryRepos,
	}
}

func (dl DeliveryService) GetAllDeliveries(ctx context.Context) (*dto.GetAllDeliveries, error) {

	deliveries, err := dl.deliRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var resp dto.GetAllDeliveries

	if err := mapper.BindingStruct(deliveries, &resp.Items); err != nil {
		return nil, err
	}

	return &resp, err
}

func (dl DeliveryService) CreateDelivery(ctx context.Context, deli *dto.CreateDeliveryRequest) (string, error) {

	entity := entities.Delivery{
		DeliveryName: deli.DeliveryName,
		DeliveryCode: deli.DeliveryCode,
		BaseCost:     deli.BaseCost,
		Description:  deli.Description,
		IsActive:     true,
	}

	deliId, err := dl.deliRepo.CreateDelivery(ctx, &entity)
	if err != nil {
		return "", err
	}

	return deliId, err
}

func (dl DeliveryService) UpdateDelivery(ctx context.Context, deli *dto.UpdateDeliveryRequest) error {
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

func (dl DeliveryService) UpdateStatusDelivery(ctx context.Context, deli *dto.UpdateStatusDeliveryRequest) error {
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
