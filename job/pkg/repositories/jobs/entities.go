package jobs

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	ID          primitive.ObjectID  `bson:"_id"`
	Name        string              `bson:"name"`
	CompanyID   primitive.ObjectID  `bson:"company_id"`
	OpenTime    primitive.DateTime  `bson:"open_time"`
	CloseTime   *primitive.DateTime `bson:"close_time,omitempty"`
	Description JobDescription      `bson:"description"`
}

func (j *Job) CreateTime() time.Time {
	return j.ID.Timestamp()
}

type JobDescription struct {
	Location    string `bson:"location"`
	Description string `bson:"description"`
}
