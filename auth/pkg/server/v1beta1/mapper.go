package v1beta1

import (
	"github.com/doogeun/auth/pkg/proto"
	"github.com/doogeun/auth/pkg/repositories/companies"
	pb "github.com/doogeun/idl/gen/go/auth/v1beta1"
)

func NewCompany(c companies.Company) *pb.Company {
	return &pb.Company{
		Id:         c.ID.Hex(),
		Name:       c.Name,
		CreateTime: proto.NewTimestamp(c.CreateTime()),
	}
}
