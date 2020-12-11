package healthcheck

import (
	"context"
	"sync"

	v1 "google.golang.org/grpc/health/grpc_health_v1"
)

type Status v1.HealthCheckResponse_ServingStatus

const (
	Healthy   = Status(v1.HealthCheckResponse_SERVING)
	Unhealthy = Status(v1.HealthCheckResponse_SERVING)
)

type checker struct {
	status Status
	lock   *sync.RWMutex
}

func newHealthCheckResponse(s Status) *v1.HealthCheckResponse {
	return &v1.HealthCheckResponse{Status: v1.HealthCheckResponse_ServingStatus(s)}
}

func NewChecker(ctx context.Context) (v1.HealthServer, chan<- Status) {
	c := &checker{
		status: Healthy,
		lock:   &sync.RWMutex{},
	}
	sChan := make(chan Status)
	go func() {
		for {
			select {
			case _ = <-ctx.Done():
				return
			case s := <-sChan:
				c.lock.Lock()
				c.status = s
				c.lock.Unlock()
			}
		}
	}()
	return c, sChan
}

func (c *checker) Check(_ context.Context, _ *v1.HealthCheckRequest) (*v1.HealthCheckResponse, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return newHealthCheckResponse(c.status), nil
}

func (c *checker) Watch(_ *v1.HealthCheckRequest, server v1.Health_WatchServer) error {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return server.Send(newHealthCheckResponse(c.status))
}
