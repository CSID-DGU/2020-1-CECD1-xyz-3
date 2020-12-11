package v1beta1

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/doogeun/gateway/pkg/grpc"
	protopkg "github.com/doogeun/gateway/pkg/proto"
	jobpbv1b1 "github.com/doogeun/idl/gen/go/job/v1beta1"
)

type handler struct {
	jc     jobpbv1b1.APIServiceClient
	jobSvc grpc.Client
}

func newHandler(jobSvc grpc.Client) *handler {
	jc := jobpbv1b1.NewAPIServiceClient(jobSvc.Connection())
	return &handler{
		jc:     jc,
		jobSvc: jobSvc,
	}
}

func Apply(g *fiber.Group, jobSvc grpc.Client) {
	h := newHandler(jobSvc)
	g.Get("/companies/:id/jobs", h.ListJobs)
	g.Get("/job/:id", h.GetJob)
}

func (h *handler) sendJSON(c *fiber.Ctx, m proto.Message) error {
	data, err := protopkg.MarshalJSON(m)
	if err != nil {
		return err
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	c.SendBytes(data)
	return nil
}

func (h *handler) ListJobs(c *fiber.Ctx) {
	companyID := c.Params("id")
	ctx, err := h.jobSvc.Context(context.Background())
	if err != nil {
		c.SendStatus(http.StatusInternalServerError)
		return
	}
	res, err := h.jc.ListJobs(ctx, &jobpbv1b1.ListJobsRequest{CompanyId: companyID})
	if err != nil {
		code := protopkg.HTTPStatusFromCode(status.Code(err))
		c.Status(code)
		c.SendString(err.Error())
		return
	}
	if err := h.sendJSON(c, res); err != nil {
		c.SendStatus(http.StatusInternalServerError)
		return
	}
}

func (h *handler) GetJob(c *fiber.Ctx) {
	jobID := c.Params("id")
	ctx, err := h.jobSvc.Context(context.Background())
	if err != nil {
		c.SendStatus(http.StatusInternalServerError)
		return
	}
	res, err := h.jc.GetJob(ctx, &jobpbv1b1.GetJobRequest{Id: jobID})
	if err != nil {
		code := protopkg.HTTPStatusFromCode(status.Code(err))
		c.Status(code)
		c.SendString(err.Error())
		return
	}
	if err := h.sendJSON(c, res); err != nil {
		c.SendStatus(http.StatusInternalServerError)
		return
	}
}
