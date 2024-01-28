package repos

import (
	"context"
	"delivery-service/internal/domain/entities"
	"delivery-service/pkgs/mongodb"
	"delivery-service/pkgs/pagable"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ShippingPackageRepos struct {
	packagesCollection *mongo.Collection
}

func NewShippingPackageRepos(client *mongodb.MongoClient) *ShippingPackageRepos {
	col := client.GetDB().Collection("shipping_packages")
	return &ShippingPackageRepos{packagesCollection: col}
}

func (dr ShippingPackageRepos) GetById(ctx context.Context, packageId string) (*entities.ShippingPackage, error) {
	var entity entities.ShippingPackage
	id, _ := primitive.ObjectIDFromHex(packageId)

	err := dr.packagesCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&entity)
	if err != nil {
		return nil, err
	}

	return &entity, err
}

func (dr ShippingPackageRepos) GetAllPackage(ctx context.Context, query *pagable.Query) ([]entities.ShippingPackage, int, error) {
	filter, err := query.ConvertQueryToFilter()
	if err != nil {
		return nil, 0, err
	}

	var packages []entities.ShippingPackage

	opts := options.Find().SetLimit(int64(query.GetSize())).SetSkip(int64(query.GetOffset()))
	cursor, err := dr.packagesCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}

	if err = cursor.All(ctx, &packages); err != nil {
		return nil, 0, err
	}

	total, err := dr.packagesCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return packages, int(total), err
}

func (dr ShippingPackageRepos) CreateShippingPackages(ctx context.Context, packages *entities.ShippingPackage) (string, error) {
	packages.CreatedAt = time.Now()
	packages.UpdatedAt = time.Now()

	lastId, err := dr.packagesCollection.InsertOne(ctx, packages)
	if err != nil {
		return "", err
	}
	return lastId.InsertedID.(primitive.ObjectID).Hex(), err
}

func (dr ShippingPackageRepos) UpdateStatus(ctx context.Context, packages *entities.ShippingPackage) error {

	update := bson.D{
		{"$set", bson.D{
			{"is_active", packages.Status},
			{"update_at", time.Now()},
		}},
	}
	data, err := dr.packagesCollection.UpdateByID(ctx, packages.Id, update)
	if err != nil {
		return err
	}

	if data.ModifiedCount == 0 {
		return errors.New("not change")
	}

	return nil
}
