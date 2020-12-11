package companies

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

func (j *Company) CreateTime() time.Time {
	return j.ID.Timestamp()
}
