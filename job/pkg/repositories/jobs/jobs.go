package jobs

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/doogeun/job/pkg/database"
	"github.com/doogeun/job/pkg/repositories"
)

const (
	jobDatabase    string = "job"
	jobsCollection string = "jobs"
)

type repository struct {
	db *mongo.Database
}

type Repository interface {
	ListJobs(ctx context.Context, companyID string, limit int32, lastID *string) ([]Job, error)
	GetJob(ctx context.Context, id string) (*Job, error)
}

func New(db database.Client) Repository {
	return &repository{
		db: db.Database(jobDatabase),
	}
}

func (r *repository) ListJobs(ctx context.Context, companyID string, limit int32, lastID *string) ([]Job, error) {
	typedCompanyID, err := primitive.ObjectIDFromHex(companyID)
	if err != nil {
		return nil, repositories.ErrInvalidArgument
	}
	filter := bson.D{{Key: "company_id", Value: typedCompanyID}}
	if lastID != nil {
		typedLastID, err := primitive.ObjectIDFromHex(*lastID)
		if err != nil {
			return nil, repositories.ErrInvalidArgument
		}
		filter = append(bson.D{{Key: "_id", Value: bson.D{{Key: "$lt", Value: typedLastID}}}}, filter...)
	}
	c, err := r.db.Collection(jobsCollection).Find(ctx, filter,
		&options.FindOptions{
			Sort:  bson.D{{Key: "_id", Value: -1}},
			Limit: database.Limit(limit),
		})
	if err != nil {
		return nil, err
	}
	var jobs []Job
	if err := c.All(ctx, &jobs); err != nil {
		return nil, err
	}
	return jobs, nil
}

func (r *repository) GetJob(ctx context.Context, id string) (*Job, error) {
	typedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, repositories.ErrInvalidArgument
	}
	res := r.db.Collection(jobsCollection).FindOne(ctx, bson.D{{Key: "_id", Value: typedID}})
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repositories.ErrNotFound
		}
		return nil, err
	}
	var job Job
	if err := res.Decode(&job); err != nil {
		return nil, err
	}
	return &job, nil
}
