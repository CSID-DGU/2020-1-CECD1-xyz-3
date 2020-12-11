package v1beta1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/doogeun/auth/pkg/repositories"
	"github.com/doogeun/auth/pkg/repositories/companies"
	pb "github.com/doogeun/idl/gen/go/auth/v1beta1"
)

type server struct {
	cr companies.Repository
}

func New(cr companies.Repository) pb.APIServiceServer {
	return &server{
		cr: cr,
	}
}

func (s *server) GetCompany(ctx context.Context, req *pb.GetCompanyRequest) (*pb.GetCompanyResponse, error) {
	company, err := s.cr.GetCompany(ctx, req.Id)
	if err != nil {
		if repositories.IsInvalidArgumentErr(err) || repositories.IsNotFoundErr(err) {
			return nil, status.Errorf(codes.NotFound, "Company %s not found", req.Id)
		}
		return nil, err
	}
	return &pb.GetCompanyResponse{Company: NewCompany(*company)}, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// TODO: Implement me
	return nil, status.Error(codes.Unimplemented, "GetUser is not implemented")
}
