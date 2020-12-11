package companies

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/doogeun/auth/pkg/database"
	"github.com/doogeun/auth/pkg/repositories"
)

const (
	companiesDatabase   string = "auth"
	companiesCollection string = "companies"
)

type repository struct {
	db *mongo.Database
}

type Repository interface {
	GetCompany(ctx context.Context, id string) (*Company, error)
}

func New(db database.Client) Repository {
	return &repository{
		db: db.Database(companiesDatabase),
	}
}

func (r *repository) GetCompany(ctx context.Context, id string) (*Company, error) {
	typedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, repositories.ErrInvalidArgument
	}
	res := r.db.Collection(companiesCollection).FindOne(ctx, bson.D{{Key: "_id", Value: typedID}})
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repositories.ErrNotFound
		}
		return nil, err
	}
	var company Company
	if err := res.Decode(&company); err != nil {
		return nil, err
	}
	return &company, nil
}
