package v1beta1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authv1b1 "github.com/doogeun/idl/gen/go/auth/v1beta1"
	pb "github.com/doogeun/idl/gen/go/job/v1beta1"
	"github.com/doogeun/job/pkg/grpc"
	"github.com/doogeun/job/pkg/repositories"
	"github.com/doogeun/job/pkg/repositories/jobs"
)

type service struct {
	ac      authv1b1.APIServiceClient
	authSvc grpc.Client
	jr      jobs.Repository
}

func New(authSvc grpc.Client, jr jobs.Repository) pb.APIServiceServer {
	ac := authv1b1.NewAPIServiceClient(authSvc.Connection())
	return &service{
		ac:      ac,
		authSvc: authSvc,
		jr:      jr,
	}
}

func (s *service) ListJobs(ctx context.Context, req *pb.ListJobsRequest) (*pb.ListJobsResponse, error) {
	lastID := &req.LastJobId
	if req.LastJobId == "" {
		lastID = nil
	}
	ctx, err := s.authSvc.Context(ctx)
	if err != nil {
		return nil, err
	}
	res, err := s.ac.GetCompany(ctx, &authv1b1.GetCompanyRequest{Id: req.CompanyId})
	if err != nil {
		return nil, err
	}
	cs := NewCompanySummary(res.Company)
	jobList, err := s.jr.ListJobs(ctx, req.CompanyId, req.Limit, lastID)
	if err != nil {
		return nil, err
	}
	data := make([]*pb.JobSummary, len(jobList))
	for idx, job := range jobList {
		data[idx] = NewJobSummary(&job, cs)
	}
	return &pb.ListJobsResponse{Jobs: data}, nil
}

func (s *service) GetJob(ctx context.Context, req *pb.GetJobRequest) (*pb.GetJobResponse, error) {
	job, err := s.jr.GetJob(ctx, req.Id)
	if err != nil {
		if repositories.IsInvalidArgumentErr(err) || repositories.IsNotFoundErr(err) {
			return nil, status.Errorf(codes.NotFound, "Job %s not found", req.Id)
		}
		return nil, err
	}
	ctx, err = s.authSvc.Context(ctx)
	if err != nil {
		return nil, err
	}
	res, err := s.ac.GetCompany(ctx, &authv1b1.GetCompanyRequest{Id: job.CompanyID.Hex()})
	if err != nil {
		return nil, err
	}
	return &pb.GetJobResponse{Job: NewJob(job, NewCompanySummary(res.Company))}, nil
}
