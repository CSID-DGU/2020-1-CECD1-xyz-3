package v1beta1

import (
	"github.com/golang/protobuf/ptypes/timestamp"

	authv1beta1 "github.com/doogeun/idl/gen/go/auth/v1beta1"
	pb "github.com/doogeun/idl/gen/go/job/v1beta1"
	"github.com/doogeun/job/pkg/proto"
	"github.com/doogeun/job/pkg/repositories/jobs"
)

func NewCompanySummary(c *authv1beta1.Company) *pb.CompanySummary {
	return &pb.CompanySummary{
		Id:   c.Id,
		Name: c.Name,
	}
}

func NewJob(j *jobs.Job, cs *pb.CompanySummary) *pb.Job {
	var closeTime *timestamp.Timestamp
	if j.CloseTime != nil {
		closeTime = proto.NewTimestamp(j.CloseTime.Time())
	}
	return &pb.Job{
		Id:          j.ID.Hex(),
		Name:        j.Name,
		Company:     cs,
		CreateTime:  proto.NewTimestamp(j.CreateTime()),
		OpenTime:    proto.NewTimestamp(j.OpenTime.Time()),
		CloseTime:   closeTime,
		Description: NewJobDescription(&j.Description),
	}
}

func NewJobDescription(jd *jobs.JobDescription) *pb.JobDescription {
	return &pb.JobDescription{
		Location:    jd.Location,
		Description: jd.Description,
	}
}

func NewJobSummary(j *jobs.Job, cs *pb.CompanySummary) *pb.JobSummary {
	return &pb.JobSummary{
		Id:      j.ID.Hex(),
		Name:    j.Name,
		Company: cs,
	}
}
