package repos

import (
	"context"
	"delivery-service/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type DeliveryRepos struct {
	deliveryCollection *mongo.Collection
}

func NewDeliveryRepos(db *mongo.Database) DeliveryRepos {
	col := db.Collection("delivery")
	return DeliveryRepos{deliveryCollection: col}
}

func (dr DeliveryRepos) GetById(ctx context.Context, deliId string) (*entities.Delivery, error) {
	var entity entities.Delivery
	id, _ := primitive.ObjectIDFromHex(deliId)

	err := dr.deliveryCollection.FindOne(ctx, entities.Delivery{ID: id}).Decode(&entity)
	if err != nil {
		return nil, err
	}

	return &entity, err
}

func (dr DeliveryRepos) GetAll(ctx context.Context) ([]entities.Delivery, error) {
	var delis []entities.Delivery

	cursor, err := dr.deliveryCollection.Find(ctx, bson.D{{}}, nil)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &delis); err != nil {
		return nil, err
	}
	return delis, err
}

func (dr DeliveryRepos) CreateDelivery(ctx context.Context, deli *entities.Delivery) (string, error) {
	deli.CreateAt = time.Now()
	deli.UpdateAt = time.Now()

	lastId, err := dr.deliveryCollection.InsertOne(ctx, deli)
	if err != nil {
		return "", err
	}
	return lastId.InsertedID.(primitive.ObjectID).Hex(), err
}

func (dr DeliveryRepos) Update(ctx context.Context, deli *entities.Delivery) error {
	_, err := dr.deliveryCollection.UpdateByID(ctx, deli.ID, deli)
	if err != nil {
		return err
	}
	return nil
}
